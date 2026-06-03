// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileByBizResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFileByBizResponseBody
	GetCode() *string
	SetData(v string) *CreateFileByBizResponseBody
	GetData() *string
	SetMessage(v string) *CreateFileByBizResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFileByBizResponseBody
	GetRequestId() *string
}

type CreateFileByBizResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileByBizResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileByBizResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileByBizResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFileByBizResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFileByBizResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFileByBizResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileByBizResponseBody) SetCode(v string) *CreateFileByBizResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFileByBizResponseBody) SetData(v string) *CreateFileByBizResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFileByBizResponseBody) SetMessage(v string) *CreateFileByBizResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFileByBizResponseBody) SetRequestId(v string) *CreateFileByBizResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileByBizResponseBody) Validate() error {
	return dara.Validate(s)
}
