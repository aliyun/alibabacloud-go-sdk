// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetGrayDomainFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v []*BatchSetGrayDomainFunctionResponseBodyDomainConfigList) *BatchSetGrayDomainFunctionResponseBody
	GetDomainConfigList() []*BatchSetGrayDomainFunctionResponseBodyDomainConfigList
	SetRequestId(v string) *BatchSetGrayDomainFunctionResponseBody
	GetRequestId() *string
}

type BatchSetGrayDomainFunctionResponseBody struct {
	DomainConfigList []*BatchSetGrayDomainFunctionResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetGrayDomainFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetGrayDomainFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetGrayDomainFunctionResponseBody) GetDomainConfigList() []*BatchSetGrayDomainFunctionResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *BatchSetGrayDomainFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetGrayDomainFunctionResponseBody) SetDomainConfigList(v []*BatchSetGrayDomainFunctionResponseBodyDomainConfigList) *BatchSetGrayDomainFunctionResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *BatchSetGrayDomainFunctionResponseBody) SetRequestId(v string) *BatchSetGrayDomainFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetGrayDomainFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchSetGrayDomainFunctionResponseBodyDomainConfigList struct {
	// example:
	//
	// 1234567
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// set_resp_header
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s BatchSetGrayDomainFunctionResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s BatchSetGrayDomainFunctionResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) SetConfigId(v int64) *BatchSetGrayDomainFunctionResponseBodyDomainConfigList {
	s.ConfigId = &v
	return s
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) SetDomainName(v string) *BatchSetGrayDomainFunctionResponseBodyDomainConfigList {
	s.DomainName = &v
	return s
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) SetFunctionName(v string) *BatchSetGrayDomainFunctionResponseBodyDomainConfigList {
	s.FunctionName = &v
	return s
}

func (s *BatchSetGrayDomainFunctionResponseBodyDomainConfigList) Validate() error {
	return dara.Validate(s)
}
