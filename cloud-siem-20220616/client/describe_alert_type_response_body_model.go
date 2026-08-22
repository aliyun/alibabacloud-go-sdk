// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertTypeResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody
	GetData() []*DescribeAlertTypeResponseBodyData
	SetMessage(v string) *DescribeAlertTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertTypeResponseBody
	GetSuccess() *bool
}

type DescribeAlertTypeResponseBody struct {
	// The request status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertTypeResponseBody) GetData() []*DescribeAlertTypeResponseBodyData {
	return s.Data
}

func (s *DescribeAlertTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertTypeResponseBody) SetCode(v int32) *DescribeAlertTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetMessage(v string) *DescribeAlertTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetRequestId(v string) *DescribeAlertTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetSuccess(v bool) *DescribeAlertTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertTypeResponseBodyData struct {
	// The threat type.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The threat type category identifier.
	//
	// example:
	//
	// identity_access
	AlertTypeCategory *string `json:"AlertTypeCategory,omitempty" xml:"AlertTypeCategory,omitempty"`
	// The threat type category name in the language of the current request. Empty if no translation is available.
	//
	// example:
	//
	// Identity and Access
	AlertTypeCategoryMds *string `json:"AlertTypeCategoryMds,omitempty" xml:"AlertTypeCategoryMds,omitempty"`
	// The display order of the threat type category.
	//
	// example:
	//
	// 10
	AlertTypeCategoryOrder *int32 `json:"AlertTypeCategoryOrder,omitempty" xml:"AlertTypeCategoryOrder,omitempty"`
	// The Medusa code of the threat type.
	//
	// example:
	//
	// siem_rule_type_process_abnormal_command
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// The English name of the threat type. Empty if no translation is available.
	//
	// example:
	//
	// Unusual Logon
	AlertTypeNameEn *string `json:"AlertTypeNameEn,omitempty" xml:"AlertTypeNameEn,omitempty"`
	// The Chinese name of the threat type. Empty if no translation is available.
	//
	// example:
	//
	// 异常登录
	AlertTypeNameZh *string `json:"AlertTypeNameZh,omitempty" xml:"AlertTypeNameZh,omitempty"`
}

func (s DescribeAlertTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeCategory() *string {
	return s.AlertTypeCategory
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeCategoryMds() *string {
	return s.AlertTypeCategoryMds
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeCategoryOrder() *int32 {
	return s.AlertTypeCategoryOrder
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeMds() *string {
	return s.AlertTypeMds
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeNameEn() *string {
	return s.AlertTypeNameEn
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeNameZh() *string {
	return s.AlertTypeNameZh
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertType(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeCategory(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeCategory = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeCategoryMds(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeCategoryMds = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeCategoryOrder(v int32) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeCategoryOrder = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeMds(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeNameEn(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeNameEn = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeNameZh(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeNameZh = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
