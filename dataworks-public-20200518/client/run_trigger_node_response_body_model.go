// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTriggerNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunTriggerNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunTriggerNodeResponseBody
	GetSuccess() *bool
}

type RunTriggerNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9BA675F1-F848-4752-A6E3-92ABA0616005
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunTriggerNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTriggerNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RunTriggerNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTriggerNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunTriggerNodeResponseBody) SetRequestId(v string) *RunTriggerNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTriggerNodeResponseBody) SetSuccess(v bool) *RunTriggerNodeResponseBody {
	s.Success = &v
	return s
}

func (s *RunTriggerNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
