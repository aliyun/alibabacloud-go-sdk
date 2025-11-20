// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *WearOrgHonorResponseBodyContent) *WearOrgHonorResponseBody
	GetContent() *WearOrgHonorResponseBodyContent
	SetErrorCode(v string) *WearOrgHonorResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *WearOrgHonorResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *WearOrgHonorResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *WearOrgHonorResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *WearOrgHonorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WearOrgHonorResponseBody
	GetSuccess() *bool
}

type WearOrgHonorResponseBody struct {
	Content *WearOrgHonorResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"errorCtx,omitempty" xml:"errorCtx,omitempty"`
	ErrorMsg  *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WearOrgHonorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorResponseBody) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBody) GetContent() *WearOrgHonorResponseBodyContent {
	return s.Content
}

func (s *WearOrgHonorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WearOrgHonorResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *WearOrgHonorResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *WearOrgHonorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WearOrgHonorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WearOrgHonorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WearOrgHonorResponseBody) SetContent(v *WearOrgHonorResponseBodyContent) *WearOrgHonorResponseBody {
	s.Content = v
	return s
}

func (s *WearOrgHonorResponseBody) SetErrorCode(v string) *WearOrgHonorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WearOrgHonorResponseBody) SetErrorCtx(v map[string]interface{}) *WearOrgHonorResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *WearOrgHonorResponseBody) SetErrorMsg(v string) *WearOrgHonorResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *WearOrgHonorResponseBody) SetHttpStatusCode(v int32) *WearOrgHonorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WearOrgHonorResponseBody) SetRequestId(v string) *WearOrgHonorResponseBody {
	s.RequestId = &v
	return s
}

func (s *WearOrgHonorResponseBody) SetSuccess(v bool) *WearOrgHonorResponseBody {
	s.Success = &v
	return s
}

func (s *WearOrgHonorResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WearOrgHonorResponseBodyContent struct {
	// example:
	//
	// null
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s WearOrgHonorResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorResponseBodyContent) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *WearOrgHonorResponseBodyContent) SetData(v interface{}) *WearOrgHonorResponseBodyContent {
	s.Data = v
	return s
}

func (s *WearOrgHonorResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
