// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonLogDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetCommonLogDetailRequest
	GetFrom() *int64
	SetHiddenProcess(v bool) *GetCommonLogDetailRequest
	GetHiddenProcess() *bool
	SetLogRequestId(v string) *GetCommonLogDetailRequest
	GetLogRequestId() *string
	SetTo(v int64) *GetCommonLogDetailRequest
	GetTo() *int64
}

type GetCommonLogDetailRequest struct {
	// The start time of the time range within which the logs that you want to query were generated. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821542
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Specifies whether to hide the process of each step. Valid values:
	//
	// 	- true: hides the process and returns only the result log of each step.
	//
	// 	- false: does not hide the process and displays the start and result logs of each step.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	HiddenProcess *bool `json:"HiddenProcess,omitempty" xml:"HiddenProcess,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The end time of the time range within which the logs that you want to query were generated. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821666
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetCommonLogDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommonLogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetCommonLogDetailRequest) GetHiddenProcess() *bool {
	return s.HiddenProcess
}

func (s *GetCommonLogDetailRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *GetCommonLogDetailRequest) GetTo() *int64 {
	return s.To
}

func (s *GetCommonLogDetailRequest) SetFrom(v int64) *GetCommonLogDetailRequest {
	s.From = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetHiddenProcess(v bool) *GetCommonLogDetailRequest {
	s.HiddenProcess = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetLogRequestId(v string) *GetCommonLogDetailRequest {
	s.LogRequestId = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetTo(v int64) *GetCommonLogDetailRequest {
	s.To = &v
	return s
}

func (s *GetCommonLogDetailRequest) Validate() error {
	return dara.Validate(s)
}
