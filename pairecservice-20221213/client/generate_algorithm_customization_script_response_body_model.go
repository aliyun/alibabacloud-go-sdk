// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAlgorithmCustomizationScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogId(v string) *GenerateAlgorithmCustomizationScriptResponseBody
	GetLogId() *string
	SetOssAddress(v string) *GenerateAlgorithmCustomizationScriptResponseBody
	GetOssAddress() *string
	SetRequestId(v string) *GenerateAlgorithmCustomizationScriptResponseBody
	GetRequestId() *string
}

type GenerateAlgorithmCustomizationScriptResponseBody struct {
	// example:
	//
	// 4
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// example:
	//
	// oss:xxxx
	OssAddress *string `json:"OssAddress,omitempty" xml:"OssAddress,omitempty"`
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAlgorithmCustomizationScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAlgorithmCustomizationScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) GetLogId() *string {
	return s.LogId
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) GetOssAddress() *string {
	return s.OssAddress
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) SetLogId(v string) *GenerateAlgorithmCustomizationScriptResponseBody {
	s.LogId = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) SetOssAddress(v string) *GenerateAlgorithmCustomizationScriptResponseBody {
	s.OssAddress = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) SetRequestId(v string) *GenerateAlgorithmCustomizationScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
