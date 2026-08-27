// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodePrototypeV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFlowNodePrototypeV2ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *ListFlowNodePrototypeV2ResponseBody
	GetCode() *int64
	SetData(v *ListFlowNodePrototypeV2ResponseBodyData) *ListFlowNodePrototypeV2ResponseBody
	GetData() *ListFlowNodePrototypeV2ResponseBodyData
	SetMessage(v string) *ListFlowNodePrototypeV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlowNodePrototypeV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlowNodePrototypeV2ResponseBody
	GetSuccess() *bool
}

type ListFlowNodePrototypeV2ResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListFlowNodePrototypeV2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlowNodePrototypeV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodePrototypeV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetData() *ListFlowNodePrototypeV2ResponseBodyData {
	return s.Data
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowNodePrototypeV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetAccessDeniedDetail(v string) *ListFlowNodePrototypeV2ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetCode(v int64) *ListFlowNodePrototypeV2ResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetData(v *ListFlowNodePrototypeV2ResponseBodyData) *ListFlowNodePrototypeV2ResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetMessage(v string) *ListFlowNodePrototypeV2ResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetRequestId(v string) *ListFlowNodePrototypeV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) SetSuccess(v bool) *ListFlowNodePrototypeV2ResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlowNodePrototypeV2ResponseBodyData struct {
	// The list of request result data.
	Model []*ListFlowNodePrototypeV2ResponseBodyDataModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
}

func (s ListFlowNodePrototypeV2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodePrototypeV2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowNodePrototypeV2ResponseBodyData) GetModel() []*ListFlowNodePrototypeV2ResponseBodyDataModel {
	return s.Model
}

func (s *ListFlowNodePrototypeV2ResponseBodyData) SetModel(v []*ListFlowNodePrototypeV2ResponseBodyDataModel) *ListFlowNodePrototypeV2ResponseBodyData {
	s.Model = v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBodyData) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowNodePrototypeV2ResponseBodyDataModel struct {
	// The component prototype code.
	//
	// example:
	//
	// SendWhatsAppMessageNode
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The component group code.
	//
	// example:
	//
	// Core
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The public extension information. This is used by the frontend to display extended information of flow components. The value is a JSON string. The first parameter in the string represents the flow component title. The following section describes the other fields:
	//
	// - en: the English information of the flow component.
	//
	// - zh: the Chinese information of the flow component.
	//
	// - name: the name of the flow component.
	//
	// - remark: the remarks of the flow component.
	//
	// - order: the display order of the flow component.
	//
	// - style: the style of the flow component.
	//
	// - svg: the icon URL of the flow component.
	//
	// - icon: a deprecated field.
	//
	// - bgcolor: the background color of the icon.
	//
	// example:
	//
	// {\\"i18n\\": {\\"en\\": {\\"name\\": \\"Send a WhatsApp Message\\", \\"remark\\": \\"Send a message with the ability for the user to reply utilizing WhatsApp specific features.\\"}, \\"zh\\": {\\"name\\": \\"Send WhatsApp messages\\", \\"remark\\": \\"Send a message that allows users to reply using specific features of WhatsApp\\"}}, \\"order\\": \\"9000\\", \\"style\\": {\\"svg\\": \\"https://img.alicdn.com/***********************************\\", \\"icon\\": \\"https://img.alicdn.com/***********************************\\", \\"bgcolor\\": \\"blue\\"}}
	PublicExtend *string `json:"PublicExtend,omitempty" xml:"PublicExtend,omitempty"`
	// The component prototype status. Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowNodePrototypeV2ResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodePrototypeV2ResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) GetCode() *string {
	return s.Code
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) GetPublicExtend() *string {
	return s.PublicExtend
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) GetStatus() *string {
	return s.Status
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) SetCode(v string) *ListFlowNodePrototypeV2ResponseBodyDataModel {
	s.Code = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) SetGroupCode(v string) *ListFlowNodePrototypeV2ResponseBodyDataModel {
	s.GroupCode = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) SetPublicExtend(v string) *ListFlowNodePrototypeV2ResponseBodyDataModel {
	s.PublicExtend = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) SetStatus(v string) *ListFlowNodePrototypeV2ResponseBodyDataModel {
	s.Status = &v
	return s
}

func (s *ListFlowNodePrototypeV2ResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}
