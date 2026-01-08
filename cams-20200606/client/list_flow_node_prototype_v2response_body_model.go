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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 91
	Code *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFlowNodePrototypeV2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// example:
	//
	// 示例值示例值
	PublicExtend *string `json:"PublicExtend,omitempty" xml:"PublicExtend,omitempty"`
	// example:
	//
	// 示例值示例值
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
