// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDigitalSignByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryDigitalSignByNameResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryDigitalSignByNameResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryDigitalSignByNameResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryDigitalSignByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDigitalSignByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDigitalSignByNameResponseBody
	GetSuccess() *bool
}

type QueryDigitalSignByNameResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDigitalSignByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDigitalSignByNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalSignByNameResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryDigitalSignByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDigitalSignByNameResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryDigitalSignByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDigitalSignByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDigitalSignByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDigitalSignByNameResponseBody) SetAccessDeniedDetail(v string) *QueryDigitalSignByNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetCode(v string) *QueryDigitalSignByNameResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetData(v map[string]interface{}) *QueryDigitalSignByNameResponseBody {
	s.Data = v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetMessage(v string) *QueryDigitalSignByNameResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetRequestId(v string) *QueryDigitalSignByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetSuccess(v bool) *QueryDigitalSignByNameResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) Validate() error {
	return dara.Validate(s)
}
