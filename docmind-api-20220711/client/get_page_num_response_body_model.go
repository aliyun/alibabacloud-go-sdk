// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPageNumResponseBodyData) *GetPageNumResponseBody
	GetData() *GetPageNumResponseBodyData
	SetErrorCode(v string) *GetPageNumResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPageNumResponseBody
	GetErrorMessage() *string
	SetHttpCode(v string) *GetPageNumResponseBody
	GetHttpCode() *string
	SetRequestId(v string) *GetPageNumResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPageNumResponseBody
	GetSuccess() *bool
}

type GetPageNumResponseBody struct {
	Data         *GetPageNumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *string                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId    *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPageNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageNumResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageNumResponseBody) GetData() *GetPageNumResponseBodyData {
	return s.Data
}

func (s *GetPageNumResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPageNumResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPageNumResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetPageNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageNumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPageNumResponseBody) SetData(v *GetPageNumResponseBodyData) *GetPageNumResponseBody {
	s.Data = v
	return s
}

func (s *GetPageNumResponseBody) SetErrorCode(v string) *GetPageNumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPageNumResponseBody) SetErrorMessage(v string) *GetPageNumResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPageNumResponseBody) SetHttpCode(v string) *GetPageNumResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPageNumResponseBody) SetRequestId(v string) *GetPageNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageNumResponseBody) SetSuccess(v bool) *GetPageNumResponseBody {
	s.Success = &v
	return s
}

func (s *GetPageNumResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPageNumResponseBodyData struct {
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s GetPageNumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPageNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPageNumResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetPageNumResponseBodyData) SetPageNum(v int32) *GetPageNumResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetPageNumResponseBodyData) Validate() error {
	return dara.Validate(s)
}
