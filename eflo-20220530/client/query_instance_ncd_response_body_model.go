// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceNcdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInstanceNcdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *QueryInstanceNcdResponseBody
	GetCode() *int32
	SetContent(v *QueryInstanceNcdResponseBodyContent) *QueryInstanceNcdResponseBody
	GetContent() *QueryInstanceNcdResponseBodyContent
	SetMessage(v string) *QueryInstanceNcdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstanceNcdResponseBody
	GetRequestId() *string
}

type QueryInstanceNcdResponseBody struct {
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
	Content *QueryInstanceNcdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:QueryInstanceNcd, arn=acs:eflo:cn-shenzhen:1263399219805497:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryInstanceNcdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceNcdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInstanceNcdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryInstanceNcdResponseBody) GetContent() *QueryInstanceNcdResponseBodyContent {
	return s.Content
}

func (s *QueryInstanceNcdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceNcdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceNcdResponseBody) SetAccessDeniedDetail(v string) *QueryInstanceNcdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetCode(v int32) *QueryInstanceNcdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetContent(v *QueryInstanceNcdResponseBodyContent) *QueryInstanceNcdResponseBody {
	s.Content = v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetMessage(v string) *QueryInstanceNcdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetRequestId(v string) *QueryInstanceNcdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInstanceNcdResponseBodyContent struct {
	// Instance 1ID
	//
	// example:
	//
	// lni-1235****
	InstanceId1 *string `json:"InstanceId1,omitempty" xml:"InstanceId1,omitempty"`
	// Instance 2ID
	//
	// example:
	//
	// lni-1234****
	InstanceId2 *string `json:"InstanceId2,omitempty" xml:"InstanceId2,omitempty"`
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
	// network communication distance between instances
	//
	// example:
	//
	// 1
	Ncd *int32 `json:"Ncd,omitempty" xml:"Ncd,omitempty"`
}

func (s QueryInstanceNcdResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceNcdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponseBodyContent) GetInstanceId1() *string {
	return s.InstanceId1
}

func (s *QueryInstanceNcdResponseBodyContent) GetInstanceId2() *string {
	return s.InstanceId2
}

func (s *QueryInstanceNcdResponseBodyContent) GetInstanceType() *string {
	return s.InstanceType
}

func (s *QueryInstanceNcdResponseBodyContent) GetNcd() *int32 {
	return s.Ncd
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceId1(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceId1 = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceId2(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceId2 = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceType(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetNcd(v int32) *QueryInstanceNcdResponseBodyContent {
	s.Ncd = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
