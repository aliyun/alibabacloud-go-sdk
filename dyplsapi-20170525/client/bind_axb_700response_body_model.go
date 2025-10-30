// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxb700ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAxb700ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAxb700ResponseBody
	GetCode() *string
	SetData(v *BindAxb700ResponseBodyData) *BindAxb700ResponseBody
	GetData() *BindAxb700ResponseBodyData
	SetMessage(v string) *BindAxb700ResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxb700ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindAxb700ResponseBody
	GetSuccess() *bool
}

type BindAxb700ResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAxb700ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxb700ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxb700ResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxb700ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAxb700ResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxb700ResponseBody) GetData() *BindAxb700ResponseBodyData {
	return s.Data
}

func (s *BindAxb700ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxb700ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxb700ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAxb700ResponseBody) SetAccessDeniedDetail(v string) *BindAxb700ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxb700ResponseBody) SetCode(v string) *BindAxb700ResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxb700ResponseBody) SetData(v *BindAxb700ResponseBodyData) *BindAxb700ResponseBody {
	s.Data = v
	return s
}

func (s *BindAxb700ResponseBody) SetMessage(v string) *BindAxb700ResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxb700ResponseBody) SetRequestId(v string) *BindAxb700ResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxb700ResponseBody) SetSuccess(v bool) *BindAxb700ResponseBody {
	s.Success = &v
	return s
}

func (s *BindAxb700ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxb700ResponseBodyData struct {
	// example:
	//
	// 700********0001
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// example:
	//
	// 123*******
	SubsId *int64 `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxb700ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAxb700ResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxb700ResponseBodyData) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindAxb700ResponseBodyData) GetSubsId() *int64 {
	return s.SubsId
}

func (s *BindAxb700ResponseBodyData) SetSecretNo(v string) *BindAxb700ResponseBodyData {
	s.SecretNo = &v
	return s
}

func (s *BindAxb700ResponseBodyData) SetSubsId(v int64) *BindAxb700ResponseBodyData {
	s.SubsId = &v
	return s
}

func (s *BindAxb700ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
