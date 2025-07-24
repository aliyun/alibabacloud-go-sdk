// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLivyComputeTokenResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLivyComputeTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLivyComputeTokenResponseBody
	GetRequestId() *string
}

type DeleteLivyComputeTokenResponseBody struct {
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

func (s DeleteLivyComputeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLivyComputeTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLivyComputeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivyComputeTokenResponseBody) SetCode(v string) *DeleteLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLivyComputeTokenResponseBody) SetMessage(v string) *DeleteLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLivyComputeTokenResponseBody) SetRequestId(v string) *DeleteLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivyComputeTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
