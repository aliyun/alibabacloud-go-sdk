// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesignateWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DesignateWorkersResponseBody
	GetCode() *int32
	SetMessage(v string) *DesignateWorkersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DesignateWorkersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DesignateWorkersResponseBody
	GetSuccess() *bool
}

type DesignateWorkersResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// job is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 765xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DesignateWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DesignateWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *DesignateWorkersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DesignateWorkersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DesignateWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DesignateWorkersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DesignateWorkersResponseBody) SetCode(v int32) *DesignateWorkersResponseBody {
	s.Code = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetMessage(v string) *DesignateWorkersResponseBody {
	s.Message = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetRequestId(v string) *DesignateWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetSuccess(v bool) *DesignateWorkersResponseBody {
	s.Success = &v
	return s
}

func (s *DesignateWorkersResponseBody) Validate() error {
	return dara.Validate(s)
}
