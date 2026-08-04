// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryDictionaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerCategoryDictionaryResponseBody
	GetCode() *string
	SetData(v *GetCustomerCategoryDictionaryResponseBodyData) *GetCustomerCategoryDictionaryResponseBody
	GetData() *GetCustomerCategoryDictionaryResponseBodyData
	SetMessage(v string) *GetCustomerCategoryDictionaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerCategoryDictionaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerCategoryDictionaryResponseBody
	GetSuccess() *bool
}

type GetCustomerCategoryDictionaryResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCustomerCategoryDictionaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerCategoryDictionaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryDictionaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerCategoryDictionaryResponseBody) GetData() *GetCustomerCategoryDictionaryResponseBodyData {
	return s.Data
}

func (s *GetCustomerCategoryDictionaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerCategoryDictionaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerCategoryDictionaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerCategoryDictionaryResponseBody) SetCode(v string) *GetCustomerCategoryDictionaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBody) SetData(v *GetCustomerCategoryDictionaryResponseBodyData) *GetCustomerCategoryDictionaryResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBody) SetMessage(v string) *GetCustomerCategoryDictionaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBody) SetRequestId(v string) *GetCustomerCategoryDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBody) SetSuccess(v bool) *GetCustomerCategoryDictionaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerCategoryDictionaryResponseBodyData struct {
	EnumConfig []*GetCustomerCategoryDictionaryResponseBodyDataEnumConfig `json:"EnumConfig,omitempty" xml:"EnumConfig,omitempty" type:"Repeated"`
}

func (s GetCustomerCategoryDictionaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryDictionaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryDictionaryResponseBodyData) GetEnumConfig() []*GetCustomerCategoryDictionaryResponseBodyDataEnumConfig {
	return s.EnumConfig
}

func (s *GetCustomerCategoryDictionaryResponseBodyData) SetEnumConfig(v []*GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) *GetCustomerCategoryDictionaryResponseBodyData {
	s.EnumConfig = v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBodyData) Validate() error {
	if s.EnumConfig != nil {
		for _, item := range s.EnumConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomerCategoryDictionaryResponseBodyDataEnumConfig struct {
	EnumName  *string `json:"enumName,omitempty" xml:"enumName,omitempty"`
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) GetEnumName() *string {
	return s.EnumName
}

func (s *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) GetEnumValue() *string {
	return s.EnumValue
}

func (s *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) SetEnumName(v string) *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig {
	s.EnumName = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) SetEnumValue(v string) *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig {
	s.EnumValue = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponseBodyDataEnumConfig) Validate() error {
	return dara.Validate(s)
}
