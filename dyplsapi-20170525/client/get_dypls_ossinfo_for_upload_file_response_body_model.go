// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDyplsOSSInfoForUploadFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetDyplsOSSInfoForUploadFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetDyplsOSSInfoForUploadFileResponseBody
	GetCode() *string
	SetData(v *GetDyplsOSSInfoForUploadFileResponseBodyData) *GetDyplsOSSInfoForUploadFileResponseBody
	GetData() *GetDyplsOSSInfoForUploadFileResponseBodyData
	SetMessage(v string) *GetDyplsOSSInfoForUploadFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDyplsOSSInfoForUploadFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDyplsOSSInfoForUploadFileResponseBody
	GetSuccess() *bool
}

type GetDyplsOSSInfoForUploadFileResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDyplsOSSInfoForUploadFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E2FD3B2F-5028-16E3-9F83-2F76F99B3873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetData() *GetDyplsOSSInfoForUploadFileResponseBodyData {
	return s.Data
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetAccessDeniedDetail(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetCode(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetData(v *GetDyplsOSSInfoForUploadFileResponseBodyData) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Data = v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetMessage(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetRequestId(v string) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) SetSuccess(v bool) *GetDyplsOSSInfoForUploadFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDyplsOSSInfoForUploadFileResponseBodyData struct {
	// example:
	//
	// LTAI***pSvPz
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 1744613007
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// https://alicom-**********-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// IjoiMjAyN*****9udGV
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// BXwW**********aoZH
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// 123456
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) GetStartPath() *string {
	return s.StartPath
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetAccessKeyId(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetExpireTime(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetHost(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetPolicy(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetSignature(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) SetStartPath(v string) *GetDyplsOSSInfoForUploadFileResponseBodyData {
	s.StartPath = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
