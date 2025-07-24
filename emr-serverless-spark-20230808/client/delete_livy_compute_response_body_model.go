// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLivyComputeResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLivyComputeResponseBody
	GetRequestId() *string
}

type DeleteLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivyComputeResponseBody) SetCode(v string) *DeleteLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLivyComputeResponseBody) SetMessage(v string) *DeleteLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLivyComputeResponseBody) SetRequestId(v string) *DeleteLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivyComputeResponseBody) Validate() error {
	return dara.Validate(s)
}
