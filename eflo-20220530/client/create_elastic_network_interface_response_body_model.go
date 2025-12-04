// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v *CreateElasticNetworkInterfaceResponseBodyContent) *CreateElasticNetworkInterfaceResponseBody
	GetContent() *CreateElasticNetworkInterfaceResponseBodyContent
	SetMessage(v string) *CreateElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type CreateElasticNetworkInterfaceResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
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
	// The response data.
	Content *CreateElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateElasticNetworkInterfaceResponseBody) GetContent() *CreateElasticNetworkInterfaceResponseBodyContent {
	return s.Content
}

func (s *CreateElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetCode(v int32) *CreateElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetContent(v *CreateElasticNetworkInterfaceResponseBodyContent) *CreateElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetMessage(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetRequestId(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateElasticNetworkInterfaceResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1fejojjo****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *CreateElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *CreateElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
