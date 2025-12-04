// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v *UpdateElasticNetworkInterfaceResponseBodyContent) *UpdateElasticNetworkInterfaceResponseBody
	GetContent() *UpdateElasticNetworkInterfaceResponseBodyContent
	SetMessage(v string) *UpdateElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type UpdateElasticNetworkInterfaceResponseBody struct {
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
	Content *UpdateElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateElasticNetworkInterfaceResponseBody) GetContent() *UpdateElasticNetworkInterfaceResponseBodyContent {
	return s.Content
}

func (s *UpdateElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetCode(v int32) *UpdateElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetContent(v *UpdateElasticNetworkInterfaceResponseBodyContent) *UpdateElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetMessage(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetRequestId(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateElasticNetworkInterfaceResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Node ID
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UpdateElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *UpdateElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
