// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteWorkerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteWorkerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkerResponseBody
	GetSuccess() *bool
}

type DeleteWorkerResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWorkerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkerResponseBody) SetCode(v string) *DeleteWorkerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkerResponseBody) SetHttpStatusCode(v int32) *DeleteWorkerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWorkerResponseBody) SetMessage(v string) *DeleteWorkerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkerResponseBody) SetRequestId(v string) *DeleteWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkerResponseBody) SetSuccess(v bool) *DeleteWorkerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkerResponseBody) Validate() error {
	return dara.Validate(s)
}
