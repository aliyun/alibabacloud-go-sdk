// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesByNcdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstancesByNcdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListInstancesByNcdResponseBody
	GetCode() *int32
	SetContent(v *ListInstancesByNcdResponseBodyContent) *ListInstancesByNcdResponseBody
	GetContent() *ListInstancesByNcdResponseBodyContent
	SetMessage(v string) *ListInstancesByNcdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesByNcdResponseBody
	GetRequestId() *string
}

type ListInstancesByNcdResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListInstancesByNcdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:ListInstancesByNcd, arn=acs:eflo:cn-heyuan:1263399219805497:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesByNcdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesByNcdResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstancesByNcdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListInstancesByNcdResponseBody) GetContent() *ListInstancesByNcdResponseBodyContent {
	return s.Content
}

func (s *ListInstancesByNcdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesByNcdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesByNcdResponseBody) SetAccessDeniedDetail(v string) *ListInstancesByNcdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetCode(v int32) *ListInstancesByNcdResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetContent(v *ListInstancesByNcdResponseBodyContent) *ListInstancesByNcdResponseBody {
	s.Content = v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetMessage(v string) *ListInstancesByNcdResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetRequestId(v string) *ListInstancesByNcdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesByNcdResponseBodyContent struct {
	// A collection of instances whose network communication distance from the source instance ID does not exceed maxNcd
	InstanceInfos []*ListInstancesByNcdResponseBodyContentInstanceInfos `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" type:"Repeated"`
	// Instance Type
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Maximum communication distance between nodes
	//
	// example:
	//
	// 3
	MaxNcd *int32 `json:"MaxNcd,omitempty" xml:"MaxNcd,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// lni-1234****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ListInstancesByNcdResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesByNcdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBodyContent) GetInstanceInfos() []*ListInstancesByNcdResponseBodyContentInstanceInfos {
	return s.InstanceInfos
}

func (s *ListInstancesByNcdResponseBodyContent) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesByNcdResponseBodyContent) GetMaxNcd() *int32 {
	return s.MaxNcd
}

func (s *ListInstancesByNcdResponseBodyContent) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListInstancesByNcdResponseBodyContent) SetInstanceInfos(v []*ListInstancesByNcdResponseBodyContentInstanceInfos) *ListInstancesByNcdResponseBodyContent {
	s.InstanceInfos = v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetInstanceType(v string) *ListInstancesByNcdResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetMaxNcd(v int32) *ListInstancesByNcdResponseBodyContent {
	s.MaxNcd = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetSourceInstanceId(v string) *ListInstancesByNcdResponseBodyContent {
	s.SourceInstanceId = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) Validate() error {
	if s.InstanceInfos != nil {
		for _, item := range s.InstanceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesByNcdResponseBodyContentInstanceInfos struct {
	// The instance ID.
	//
	// example:
	//
	// lni-1235****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// network communication distance
	//
	// example:
	//
	// 2
	Ncd *int32 `json:"Ncd,omitempty" xml:"Ncd,omitempty"`
}

func (s ListInstancesByNcdResponseBodyContentInstanceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesByNcdResponseBodyContentInstanceInfos) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) GetNcd() *int32 {
	return s.Ncd
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) SetInstanceId(v string) *ListInstancesByNcdResponseBodyContentInstanceInfos {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) SetNcd(v int32) *ListInstancesByNcdResponseBodyContentInstanceInfos {
	s.Ncd = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) Validate() error {
	return dara.Validate(s)
}
