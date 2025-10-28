// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertK8sResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ConvertK8sResourceResponseBody
	GetCode() *int32
	SetMessage(v string) *ConvertK8sResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertK8sResourceResponseBody
	GetRequestId() *string
}

type ConvertK8sResourceResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// convert success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BA938591-*********-9690-BFD3F4DD7A93
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertK8sResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertK8sResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ConvertK8sResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertK8sResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertK8sResourceResponseBody) SetCode(v int32) *ConvertK8sResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertK8sResourceResponseBody) SetMessage(v string) *ConvertK8sResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertK8sResourceResponseBody) SetRequestId(v string) *ConvertK8sResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertK8sResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
