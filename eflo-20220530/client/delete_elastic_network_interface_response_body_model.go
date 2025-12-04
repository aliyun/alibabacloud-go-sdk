// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v *DeleteElasticNetworkInterfaceResponseBodyContent) *DeleteElasticNetworkInterfaceResponseBody
	GetContent() *DeleteElasticNetworkInterfaceResponseBodyContent
	SetMessage(v string) *DeleteElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type DeleteElasticNetworkInterfaceResponseBody struct {
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
	Content *DeleteElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteElasticNetworkInterfaceResponseBody) GetContent() *DeleteElasticNetworkInterfaceResponseBodyContent {
	return s.Content
}

func (s *DeleteElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetCode(v int32) *DeleteElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetContent(v *DeleteElasticNetworkInterfaceResponseBodyContent) *DeleteElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetMessage(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetRequestId(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteElasticNetworkInterfaceResponseBodyContent struct {
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

func (s DeleteElasticNetworkInterfaceResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *DeleteElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *DeleteElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
