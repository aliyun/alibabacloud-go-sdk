// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGxb700ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindGxb700ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindGxb700ResponseBody
	GetCode() *string
	SetData(v *BindGxb700ResponseBodyData) *BindGxb700ResponseBody
	GetData() *BindGxb700ResponseBodyData
	SetMessage(v string) *BindGxb700ResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindGxb700ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindGxb700ResponseBody
	GetSuccess() *bool
}

type BindGxb700ResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. Valid values:
	//
	// OK: The request was successful.
	//
	// Other values indicate failure.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *BindGxb700ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// true: Successful.
	//
	// false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindGxb700ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindGxb700ResponseBody) GoString() string {
	return s.String()
}

func (s *BindGxb700ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindGxb700ResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindGxb700ResponseBody) GetData() *BindGxb700ResponseBodyData {
	return s.Data
}

func (s *BindGxb700ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindGxb700ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindGxb700ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindGxb700ResponseBody) SetAccessDeniedDetail(v string) *BindGxb700ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindGxb700ResponseBody) SetCode(v string) *BindGxb700ResponseBody {
	s.Code = &v
	return s
}

func (s *BindGxb700ResponseBody) SetData(v *BindGxb700ResponseBodyData) *BindGxb700ResponseBody {
	s.Data = v
	return s
}

func (s *BindGxb700ResponseBody) SetMessage(v string) *BindGxb700ResponseBody {
	s.Message = &v
	return s
}

func (s *BindGxb700ResponseBody) SetRequestId(v string) *BindGxb700ResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindGxb700ResponseBody) SetSuccess(v bool) *BindGxb700ResponseBody {
	s.Success = &v
	return s
}

func (s *BindGxb700ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindGxb700ResponseBodyData struct {
	// The purchased 700 private number, which is the X number.
	//
	// example:
	//
	// 700********0001
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The number binding ID.
	//
	// example:
	//
	// 123*******
	SubsId *int64 `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindGxb700ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindGxb700ResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindGxb700ResponseBodyData) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindGxb700ResponseBodyData) GetSubsId() *int64 {
	return s.SubsId
}

func (s *BindGxb700ResponseBodyData) SetSecretNo(v string) *BindGxb700ResponseBodyData {
	s.SecretNo = &v
	return s
}

func (s *BindGxb700ResponseBodyData) SetSubsId(v int64) *BindGxb700ResponseBodyData {
	s.SubsId = &v
	return s
}

func (s *BindGxb700ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
