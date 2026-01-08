// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFlowNodeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *ListFlowNodeGroupResponseBody
	GetCode() *int64
	SetData(v *ListFlowNodeGroupResponseBodyData) *ListFlowNodeGroupResponseBody
	GetData() *ListFlowNodeGroupResponseBodyData
	SetMessage(v string) *ListFlowNodeGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlowNodeGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlowNodeGroupResponseBody
	GetSuccess() *bool
}

type ListFlowNodeGroupResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *int64                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFlowNodeGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlowNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFlowNodeGroupResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListFlowNodeGroupResponseBody) GetData() *ListFlowNodeGroupResponseBodyData {
	return s.Data
}

func (s *ListFlowNodeGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlowNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowNodeGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlowNodeGroupResponseBody) SetAccessDeniedDetail(v string) *ListFlowNodeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFlowNodeGroupResponseBody) SetCode(v int64) *ListFlowNodeGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowNodeGroupResponseBody) SetData(v *ListFlowNodeGroupResponseBodyData) *ListFlowNodeGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowNodeGroupResponseBody) SetMessage(v string) *ListFlowNodeGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowNodeGroupResponseBody) SetRequestId(v string) *ListFlowNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodeGroupResponseBody) SetSuccess(v bool) *ListFlowNodeGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlowNodeGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlowNodeGroupResponseBodyData struct {
	Model []*ListFlowNodeGroupResponseBodyDataModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
}

func (s ListFlowNodeGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodeGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowNodeGroupResponseBodyData) GetModel() []*ListFlowNodeGroupResponseBodyDataModel {
	return s.Model
}

func (s *ListFlowNodeGroupResponseBodyData) SetModel(v []*ListFlowNodeGroupResponseBodyDataModel) *ListFlowNodeGroupResponseBodyData {
	s.Model = v
	return s
}

func (s *ListFlowNodeGroupResponseBodyData) Validate() error {
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

type ListFlowNodeGroupResponseBodyDataModel struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	PublicExtend *string `json:"PublicExtend,omitempty" xml:"PublicExtend,omitempty"`
}

func (s ListFlowNodeGroupResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodeGroupResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *ListFlowNodeGroupResponseBodyDataModel) GetCode() *string {
	return s.Code
}

func (s *ListFlowNodeGroupResponseBodyDataModel) GetPublicExtend() *string {
	return s.PublicExtend
}

func (s *ListFlowNodeGroupResponseBodyDataModel) SetCode(v string) *ListFlowNodeGroupResponseBodyDataModel {
	s.Code = &v
	return s
}

func (s *ListFlowNodeGroupResponseBodyDataModel) SetPublicExtend(v string) *ListFlowNodeGroupResponseBodyDataModel {
	s.PublicExtend = &v
	return s
}

func (s *ListFlowNodeGroupResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}
