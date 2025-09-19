// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmVirusEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationAll(v int32) *ConfirmVirusEventsRequest
	GetOperationAll() *int32
	SetOperationCode(v string) *ConfirmVirusEventsRequest
	GetOperationCode() *string
	SetOperationRange(v string) *ConfirmVirusEventsRequest
	GetOperationRange() *string
}

type ConfirmVirusEventsRequest struct {
	// Specifies whether to handle all alert events. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperationAll *int32 `json:"OperationAll,omitempty" xml:"OperationAll,omitempty"`
	// The operation that you want to perform on the alert events. Valid values:
	//
	// 	- **default**: performs in-depth detection and removal
	//
	// 	- **ignore**: ignores the alert event
	//
	// 	- **advance_mark_mis_info**: adds the alert events to the whitelist
	//
	// 	- **manual_handled**: marks the alert events as manually handled
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The server on which you want to perform the alert events.
	//
	// example:
	//
	// [{\\"type\\":\\"machine\\",\\"list\\":[\\"3aedba3d-bd4d-4dfb-bb0d-xxxxxxxxxxxx\\"]}]
	OperationRange *string `json:"OperationRange,omitempty" xml:"OperationRange,omitempty"`
}

func (s ConfirmVirusEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmVirusEventsRequest) GoString() string {
	return s.String()
}

func (s *ConfirmVirusEventsRequest) GetOperationAll() *int32 {
	return s.OperationAll
}

func (s *ConfirmVirusEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *ConfirmVirusEventsRequest) GetOperationRange() *string {
	return s.OperationRange
}

func (s *ConfirmVirusEventsRequest) SetOperationAll(v int32) *ConfirmVirusEventsRequest {
	s.OperationAll = &v
	return s
}

func (s *ConfirmVirusEventsRequest) SetOperationCode(v string) *ConfirmVirusEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *ConfirmVirusEventsRequest) SetOperationRange(v string) *ConfirmVirusEventsRequest {
	s.OperationRange = &v
	return s
}

func (s *ConfirmVirusEventsRequest) Validate() error {
	return dara.Validate(s)
}
