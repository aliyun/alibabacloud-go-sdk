// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigXResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ConfigXResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ConfigXResponseBody
	GetCode() *string
	SetData(v *ConfigXResponseBodyData) *ConfigXResponseBody
	GetData() *ConfigXResponseBodyData
	SetMessage(v string) *ConfigXResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigXResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigXResponseBody
	GetSuccess() *bool
}

type ConfigXResponseBody struct {
	// The details of the access denied error.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned status code. Valid values:
	//
	// - 0000: Success.
	//
	// - Other values: Failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration result of the work number.
	//
	// example:
	//
	// {"Code":"0000","Message":"Success","Success":true}
	Data *ConfigXResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A2568E1-7D83-5491-8DEF-600D8AFF97E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: Success.
	//
	//
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigXResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ConfigXResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigXResponseBody) GetData() *ConfigXResponseBodyData {
	return s.Data
}

func (s *ConfigXResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigXResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigXResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigXResponseBody) SetAccessDeniedDetail(v string) *ConfigXResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ConfigXResponseBody) SetCode(v string) *ConfigXResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBody) SetData(v *ConfigXResponseBodyData) *ConfigXResponseBody {
	s.Data = v
	return s
}

func (s *ConfigXResponseBody) SetMessage(v string) *ConfigXResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBody) SetRequestId(v string) *ConfigXResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigXResponseBody) SetSuccess(v bool) *ConfigXResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigXResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigXResponseBodyData struct {
	// The returned status code. A value of 0000 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigXResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ConfigXResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ConfigXResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigXResponseBodyData) SetCode(v string) *ConfigXResponseBodyData {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBodyData) SetMessage(v string) *ConfigXResponseBodyData {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBodyData) SetSuccess(v bool) *ConfigXResponseBodyData {
	s.Success = &v
	return s
}

func (s *ConfigXResponseBodyData) Validate() error {
	return dara.Validate(s)
}
