// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateVhostResponseBody
	GetCode() *int32
	SetData(v bool) *CreateVhostResponseBody
	GetData() *bool
	SetMessage(v string) *CreateVhostResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVhostResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVhostResponseBody
	GetSuccess() *bool
}

type CreateVhostResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVhostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVhostResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVhostResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateVhostResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVhostResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVhostResponseBody) SetCode(v int32) *CreateVhostResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVhostResponseBody) SetData(v bool) *CreateVhostResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVhostResponseBody) SetMessage(v string) *CreateVhostResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVhostResponseBody) SetRequestId(v string) *CreateVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVhostResponseBody) SetSuccess(v bool) *CreateVhostResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVhostResponseBody) Validate() error {
	return dara.Validate(s)
}
