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
	// Specifies whether to handle all alerts. Valid values:
	//
	// - **1**: Yes.
	//
	// - **0**: No.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperationAll *int32 `json:"OperationAll,omitempty" xml:"OperationAll,omitempty"`
	// The method to handle the alert event. Valid values:
	//
	// - **default**: deep scan and removal
	//
	// - **ignore**: ignore
	//
	// - **advance_mark_mis_info**: add to whitelist
	//
	// - **manual_handled**: manually handled
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The scope of the trojan scan alert handling. This parameter is required when OperationAll is set to 0. This parameter is ignored when OperationAll is set to 1.
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
