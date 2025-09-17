// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopAppResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *StopAppResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopAppResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StopAppResourcesResponseBody
	GetSuccess() *string
}

type StopAppResourcesResponseBody struct {
	// example:
	//
	// InvalidAppInstanceGroup.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The app instance group is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 829444D6-9FD3-5C65-A570-065975537647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopAppResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopAppResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopAppResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAppResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StopAppResourcesResponseBody) SetCode(v string) *StopAppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *StopAppResourcesResponseBody) SetMessage(v string) *StopAppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *StopAppResourcesResponseBody) SetRequestId(v string) *StopAppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAppResourcesResponseBody) SetSuccess(v string) *StopAppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *StopAppResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
