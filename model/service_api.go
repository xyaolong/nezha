package model

import "time"

type ServiceForm struct {
	Name                string          `json:"name,omitempty"`
	Target              string          `json:"target,omitempty"`
	Type                uint8           `json:"type,omitempty"`
	Cover               uint8           `json:"cover,omitempty"`
	Notify              bool            `json:"notify,omitempty"`
	Duration            uint64          `json:"duration,omitempty"`
	MinLatency          float32         `json:"min_latency,omitempty"`
	MaxLatency          float32         `json:"max_latency,omitempty"`
	LatencyNotify       bool            `json:"latency_notify,omitempty"`
	EnableTriggerTask   bool            `json:"enable_trigger_task,omitempty"`
	EnableShowInService bool            `json:"enable_show_in_service,omitempty"`
	FailTriggerTasks    []uint64        `json:"fail_trigger_tasks,omitempty"`
	RecoverTriggerTasks []uint64        `json:"recover_trigger_tasks,omitempty"`
	SkipServers         map[uint64]bool `json:"skip_servers,omitempty"`
	NotificationGroupID uint64          `json:"notification_group_id,omitempty"`
}

type ServiceResponseItem struct {
	Service     *Service     `json:"service,omitempty"`
	CurrentUp   uint64       `json:"current_up"`
	CurrentDown uint64       `json:"current_down"`
	TotalUp     uint64       `json:"total_up"`
	TotalDown   uint64       `json:"total_down"`
	Delay       *[30]float32 `json:"delay,omitempty"`
	Up          *[30]int     `json:"up,omitempty"`
	Down        *[30]int     `json:"down,omitempty"`
}

func (r ServiceResponseItem) TotalUptime() float32 {
	if r.TotalUp+r.TotalDown == 0 {
		return 0
	}
	return float32(r.TotalUp) / (float32(r.TotalUp + r.TotalDown)) * 100
}

type CycleTransferStats struct {
	Name       string               `json:"name"`
	From       time.Time            `json:"from"`
	To         time.Time            `json:"to"`
	Max        uint64               `json:"max"`
	Min        uint64               `json:"min"`
	ServerName map[uint64]string    `json:"server_name,omitempty"`
	Transfer   map[uint64]uint64    `json:"transfer,omitempty"`
	NextUpdate map[uint64]time.Time `json:"next_update,omitempty"`
}

type ServiceResponse struct {
	Services           map[uint64]ServiceResponseItem `json:"services,omitempty"`
	CycleTransferStats map[uint64]CycleTransferStats  `json:"cycle_transfer_stats,omitempty"`
}
