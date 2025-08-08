// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMcdpMaterialResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcdpMaterialResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMcdpMaterialResponseBodyResultContent) *CreateMcdpMaterialResponseBody
	GetResultContent() *CreateMcdpMaterialResponseBodyResultContent
	SetResultMessage(v string) *CreateMcdpMaterialResponseBody
	GetResultMessage() *string
}

type CreateMcdpMaterialResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMcdpMaterialResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcdpMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcdpMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcdpMaterialResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcdpMaterialResponseBody) GetResultContent() *CreateMcdpMaterialResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMcdpMaterialResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcdpMaterialResponseBody) SetRequestId(v string) *CreateMcdpMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcdpMaterialResponseBody) SetResultCode(v string) *CreateMcdpMaterialResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcdpMaterialResponseBody) SetResultContent(v *CreateMcdpMaterialResponseBodyResultContent) *CreateMcdpMaterialResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMcdpMaterialResponseBody) SetResultMessage(v string) *CreateMcdpMaterialResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcdpMaterialResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcdpMaterialResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcdpMaterialResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpMaterialResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMcdpMaterialResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *CreateMcdpMaterialResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CreateMcdpMaterialResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *CreateMcdpMaterialResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcdpMaterialResponseBodyResultContent) SetCode(v string) *CreateMcdpMaterialResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *CreateMcdpMaterialResponseBodyResultContent) SetData(v string) *CreateMcdpMaterialResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CreateMcdpMaterialResponseBodyResultContent) SetMessage(v string) *CreateMcdpMaterialResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *CreateMcdpMaterialResponseBodyResultContent) SetSuccess(v bool) *CreateMcdpMaterialResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *CreateMcdpMaterialResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
