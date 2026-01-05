// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlotStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SlotStatus
	GetCode() *string
	SetDetail(v *SlotStatusDetail) *SlotStatus
	GetDetail() *SlotStatusDetail
	SetMessage(v string) *SlotStatus
	GetMessage() *string
	SetPhase(v string) *SlotStatus
	GetPhase() *string
}

type SlotStatus struct {
	// example:
	//
	// 200
	Code   *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Detail *SlotStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// Init Succeed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
}

func (s SlotStatus) String() string {
	return dara.Prettify(s)
}

func (s SlotStatus) GoString() string {
	return s.String()
}

func (s *SlotStatus) GetCode() *string {
	return s.Code
}

func (s *SlotStatus) GetDetail() *SlotStatusDetail {
	return s.Detail
}

func (s *SlotStatus) GetMessage() *string {
	return s.Message
}

func (s *SlotStatus) GetPhase() *string {
	return s.Phase
}

func (s *SlotStatus) SetCode(v string) *SlotStatus {
	s.Code = &v
	return s
}

func (s *SlotStatus) SetDetail(v *SlotStatusDetail) *SlotStatus {
	s.Detail = v
	return s
}

func (s *SlotStatus) SetMessage(v string) *SlotStatus {
	s.Message = &v
	return s
}

func (s *SlotStatus) SetPhase(v string) *SlotStatus {
	s.Phase = &v
	return s
}

func (s *SlotStatus) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
