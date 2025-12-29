// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPackageTypeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPackageTypeInfoResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryPackageTypeInfoResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryPackageTypeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPackageTypeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPackageTypeInfoResponseBody
	GetSuccess() *bool
}

type QueryPackageTypeInfoResponseBody struct {
	// example:
	//
	// 示例值示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPackageTypeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPackageTypeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPackageTypeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPackageTypeInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryPackageTypeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPackageTypeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPackageTypeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPackageTypeInfoResponseBody) SetCode(v string) *QueryPackageTypeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetData(v map[string]interface{}) *QueryPackageTypeInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetMessage(v string) *QueryPackageTypeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetRequestId(v string) *QueryPackageTypeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetSuccess(v bool) *QueryPackageTypeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
