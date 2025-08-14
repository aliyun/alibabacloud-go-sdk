// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetEventCallbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetEventCallbackResponseBody
	GetSuccess() *bool
}

type SetEventCallbackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the configuration was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetEventCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetEventCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SetEventCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetEventCallbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetEventCallbackResponseBody) SetRequestId(v string) *SetEventCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEventCallbackResponseBody) SetSuccess(v bool) *SetEventCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *SetEventCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
