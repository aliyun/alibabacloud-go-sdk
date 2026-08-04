// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunIdByPkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunId(v string) *GetAliyunIdByPkResponseBody
	GetAliyunId() *string
	SetCode(v string) *GetAliyunIdByPkResponseBody
	GetCode() *string
	SetMessage(v string) *GetAliyunIdByPkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAliyunIdByPkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAliyunIdByPkResponseBody
	GetSuccess() *bool
}

type GetAliyunIdByPkResponseBody struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAliyunIdByPkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunIdByPkResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliyunIdByPkResponseBody) GetAliyunId() *string {
	return s.AliyunId
}

func (s *GetAliyunIdByPkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAliyunIdByPkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAliyunIdByPkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAliyunIdByPkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAliyunIdByPkResponseBody) SetAliyunId(v string) *GetAliyunIdByPkResponseBody {
	s.AliyunId = &v
	return s
}

func (s *GetAliyunIdByPkResponseBody) SetCode(v string) *GetAliyunIdByPkResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliyunIdByPkResponseBody) SetMessage(v string) *GetAliyunIdByPkResponseBody {
	s.Message = &v
	return s
}

func (s *GetAliyunIdByPkResponseBody) SetRequestId(v string) *GetAliyunIdByPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliyunIdByPkResponseBody) SetSuccess(v bool) *GetAliyunIdByPkResponseBody {
	s.Success = &v
	return s
}

func (s *GetAliyunIdByPkResponseBody) Validate() error {
	return dara.Validate(s)
}
