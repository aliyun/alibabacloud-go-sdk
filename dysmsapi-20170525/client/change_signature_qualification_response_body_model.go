// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSignatureQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChangeSignatureQualificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChangeSignatureQualificationResponseBody
	GetCode() *string
	SetData(v *ChangeSignatureQualificationResponseBodyData) *ChangeSignatureQualificationResponseBody
	GetData() *ChangeSignatureQualificationResponseBodyData
	SetMessage(v string) *ChangeSignatureQualificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeSignatureQualificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeSignatureQualificationResponseBody
	GetSuccess() *bool
}

type ChangeSignatureQualificationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ChangeSignatureQualificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0A974B78-02BF-4C79-ADF3-90CFBA1B55B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeSignatureQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeSignatureQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeSignatureQualificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChangeSignatureQualificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeSignatureQualificationResponseBody) GetData() *ChangeSignatureQualificationResponseBodyData {
	return s.Data
}

func (s *ChangeSignatureQualificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeSignatureQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeSignatureQualificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeSignatureQualificationResponseBody) SetAccessDeniedDetail(v string) *ChangeSignatureQualificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) SetCode(v string) *ChangeSignatureQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) SetData(v *ChangeSignatureQualificationResponseBodyData) *ChangeSignatureQualificationResponseBody {
	s.Data = v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) SetMessage(v string) *ChangeSignatureQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) SetRequestId(v string) *ChangeSignatureQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) SetSuccess(v bool) *ChangeSignatureQualificationResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeSignatureQualificationResponseBodyData struct {
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 示例值示例值
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeSignatureQualificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeSignatureQualificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeSignatureQualificationResponseBodyData) GetData() map[string]interface{} {
	return s.Data
}

func (s *ChangeSignatureQualificationResponseBodyData) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChangeSignatureQualificationResponseBodyData) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChangeSignatureQualificationResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeSignatureQualificationResponseBodyData) SetData(v map[string]interface{}) *ChangeSignatureQualificationResponseBodyData {
	s.Data = v
	return s
}

func (s *ChangeSignatureQualificationResponseBodyData) SetErrCode(v string) *ChangeSignatureQualificationResponseBodyData {
	s.ErrCode = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBodyData) SetErrMessage(v string) *ChangeSignatureQualificationResponseBodyData {
	s.ErrMessage = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBodyData) SetSuccess(v bool) *ChangeSignatureQualificationResponseBodyData {
	s.Success = &v
	return s
}

func (s *ChangeSignatureQualificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
