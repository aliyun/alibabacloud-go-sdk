// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpListLogResult interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*PdpLog) *PdpListLogResult
	GetLogs() []*PdpLog
}

type PdpListLogResult struct {
	Logs []*PdpLog `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
}

func (s PdpListLogResult) String() string {
	return dara.Prettify(s)
}

func (s PdpListLogResult) GoString() string {
	return s.String()
}

func (s *PdpListLogResult) GetLogs() []*PdpLog {
	return s.Logs
}

func (s *PdpListLogResult) SetLogs(v []*PdpLog) *PdpListLogResult {
	s.Logs = v
	return s
}

func (s *PdpListLogResult) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
