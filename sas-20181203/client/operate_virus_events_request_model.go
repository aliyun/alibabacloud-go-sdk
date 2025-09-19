// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVirusEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationAll(v int32) *OperateVirusEventsRequest
	GetOperationAll() *int32
	SetOperationCode(v string) *OperateVirusEventsRequest
	GetOperationCode() *string
	SetOperationRange(v string) *OperateVirusEventsRequest
	GetOperationRange() *string
}

type OperateVirusEventsRequest struct {
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
	// The handling scope.
	//
	// example:
	//
	// [{\\"type\\":\\"machine\\",\\"list\\":[\\"xxxxxxxxx-4cbf-4ca6-a1b7-8a09d1f86ab0\\"]}]
	OperationRange *string `json:"OperationRange,omitempty" xml:"OperationRange,omitempty"`
}

func (s OperateVirusEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateVirusEventsRequest) GoString() string {
	return s.String()
}

func (s *OperateVirusEventsRequest) GetOperationAll() *int32 {
	return s.OperationAll
}

func (s *OperateVirusEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *OperateVirusEventsRequest) GetOperationRange() *string {
	return s.OperationRange
}

func (s *OperateVirusEventsRequest) SetOperationAll(v int32) *OperateVirusEventsRequest {
	s.OperationAll = &v
	return s
}

func (s *OperateVirusEventsRequest) SetOperationCode(v string) *OperateVirusEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *OperateVirusEventsRequest) SetOperationRange(v string) *OperateVirusEventsRequest {
	s.OperationRange = &v
	return s
}

func (s *OperateVirusEventsRequest) Validate() error {
	return dara.Validate(s)
}
