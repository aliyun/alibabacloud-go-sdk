// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteK8sConfigMapResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteK8sConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteK8sConfigMapResponseBody
	GetRequestId() *string
}

type DeleteK8sConfigMapResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteK8sConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sConfigMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteK8sConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteK8sConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteK8sConfigMapResponseBody) SetCode(v int32) *DeleteK8sConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteK8sConfigMapResponseBody) SetMessage(v string) *DeleteK8sConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteK8sConfigMapResponseBody) SetRequestId(v string) *DeleteK8sConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteK8sConfigMapResponseBody) Validate() error {
	return dara.Validate(s)
}
