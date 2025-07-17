// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSyntheticTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchSyntheticTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *SwitchSyntheticTaskStatusResponseBody
	GetResult() *string
}

type SwitchSyntheticTaskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SwitchSyntheticTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchSyntheticTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSyntheticTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchSyntheticTaskStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *SwitchSyntheticTaskStatusResponseBody) SetRequestId(v string) *SwitchSyntheticTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSyntheticTaskStatusResponseBody) SetResult(v string) *SwitchSyntheticTaskStatusResponseBody {
	s.Result = &v
	return s
}

func (s *SwitchSyntheticTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
