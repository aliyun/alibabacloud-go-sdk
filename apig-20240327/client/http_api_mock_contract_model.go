// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiMockContract interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *HttpApiMockContract
	GetEnable() *bool
	SetResponseCode(v int32) *HttpApiMockContract
	GetResponseCode() *int32
	SetResponseContent(v string) *HttpApiMockContract
	GetResponseContent() *string
}

type HttpApiMockContract struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 200
	ResponseCode    *int32  `json:"responseCode,omitempty" xml:"responseCode,omitempty"`
	ResponseContent *string `json:"responseContent,omitempty" xml:"responseContent,omitempty"`
}

func (s HttpApiMockContract) String() string {
	return dara.Prettify(s)
}

func (s HttpApiMockContract) GoString() string {
	return s.String()
}

func (s *HttpApiMockContract) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiMockContract) GetResponseCode() *int32 {
	return s.ResponseCode
}

func (s *HttpApiMockContract) GetResponseContent() *string {
	return s.ResponseContent
}

func (s *HttpApiMockContract) SetEnable(v bool) *HttpApiMockContract {
	s.Enable = &v
	return s
}

func (s *HttpApiMockContract) SetResponseCode(v int32) *HttpApiMockContract {
	s.ResponseCode = &v
	return s
}

func (s *HttpApiMockContract) SetResponseContent(v string) *HttpApiMockContract {
	s.ResponseContent = &v
	return s
}

func (s *HttpApiMockContract) Validate() error {
	return dara.Validate(s)
}
