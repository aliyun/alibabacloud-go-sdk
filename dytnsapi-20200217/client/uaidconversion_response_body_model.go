// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDConversionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UAIDConversionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UAIDConversionResponseBody
	GetCode() *string
	SetMessage(v string) *UAIDConversionResponseBody
	GetMessage() *string
	SetModel(v *UAIDConversionResponseBodyModel) *UAIDConversionResponseBody
	GetModel() *UAIDConversionResponseBodyModel
	SetRequestId(v string) *UAIDConversionResponseBody
	GetRequestId() *string
}

type UAIDConversionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UAIDConversionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDConversionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UAIDConversionResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UAIDConversionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UAIDConversionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UAIDConversionResponseBody) GetModel() *UAIDConversionResponseBodyModel {
	return s.Model
}

func (s *UAIDConversionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UAIDConversionResponseBody) SetAccessDeniedDetail(v string) *UAIDConversionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDConversionResponseBody) SetCode(v string) *UAIDConversionResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDConversionResponseBody) SetMessage(v string) *UAIDConversionResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDConversionResponseBody) SetModel(v *UAIDConversionResponseBodyModel) *UAIDConversionResponseBody {
	s.Model = v
	return s
}

func (s *UAIDConversionResponseBody) SetRequestId(v string) *UAIDConversionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UAIDConversionResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UAIDConversionResponseBodyModel struct {
	// example:
	//
	// 示例值示例值示例值
	PhoneList *string `json:"PhoneList,omitempty" xml:"PhoneList,omitempty"`
}

func (s UAIDConversionResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s UAIDConversionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponseBodyModel) GetPhoneList() *string {
	return s.PhoneList
}

func (s *UAIDConversionResponseBodyModel) SetPhoneList(v string) *UAIDConversionResponseBodyModel {
	s.PhoneList = &v
	return s
}

func (s *UAIDConversionResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
