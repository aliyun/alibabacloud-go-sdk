// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMpaasAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMpaasAppInfoResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateMpaasAppInfoResponseBody
	GetResultCode() *string
	SetResultContent(v *UpdateMpaasAppInfoResponseBodyResultContent) *UpdateMpaasAppInfoResponseBody
	GetResultContent() *UpdateMpaasAppInfoResponseBodyResultContent
	SetResultMessage(v string) *UpdateMpaasAppInfoResponseBody
	GetResultMessage() *string
}

type UpdateMpaasAppInfoResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *UpdateMpaasAppInfoResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMpaasAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMpaasAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMpaasAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMpaasAppInfoResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateMpaasAppInfoResponseBody) GetResultContent() *UpdateMpaasAppInfoResponseBodyResultContent {
	return s.ResultContent
}

func (s *UpdateMpaasAppInfoResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateMpaasAppInfoResponseBody) SetRequestId(v string) *UpdateMpaasAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBody) SetResultCode(v string) *UpdateMpaasAppInfoResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBody) SetResultContent(v *UpdateMpaasAppInfoResponseBodyResultContent) *UpdateMpaasAppInfoResponseBody {
	s.ResultContent = v
	return s
}

func (s *UpdateMpaasAppInfoResponseBody) SetResultMessage(v string) *UpdateMpaasAppInfoResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMpaasAppInfoResponseBodyResultContent struct {
	Data      *UpdateMpaasAppInfoResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMpaasAppInfoResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateMpaasAppInfoResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *UpdateMpaasAppInfoResponseBodyResultContent) GetData() *UpdateMpaasAppInfoResponseBodyResultContentData {
	return s.Data
}

func (s *UpdateMpaasAppInfoResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMpaasAppInfoResponseBodyResultContent) SetData(v *UpdateMpaasAppInfoResponseBodyResultContentData) *UpdateMpaasAppInfoResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContent) SetRequestId(v string) *UpdateMpaasAppInfoResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type UpdateMpaasAppInfoResponseBodyResultContentData struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMpaasAppInfoResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMpaasAppInfoResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) GetCode() *string {
	return s.Code
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) GetData() *string {
	return s.Data
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) GetMessage() *string {
	return s.Message
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) SetAppId(v string) *UpdateMpaasAppInfoResponseBodyResultContentData {
	s.AppId = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) SetCode(v string) *UpdateMpaasAppInfoResponseBodyResultContentData {
	s.Code = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) SetData(v string) *UpdateMpaasAppInfoResponseBodyResultContentData {
	s.Data = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) SetMessage(v string) *UpdateMpaasAppInfoResponseBodyResultContentData {
	s.Message = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) SetSuccess(v bool) *UpdateMpaasAppInfoResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *UpdateMpaasAppInfoResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
