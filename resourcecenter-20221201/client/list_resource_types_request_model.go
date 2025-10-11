// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListResourceTypesRequest
	GetAcceptLanguage() *string
	SetQuery(v []*string) *ListResourceTypesRequest
	GetQuery() []*string
	SetResourceType(v string) *ListResourceTypesRequest
	GetResourceType() *string
}

type ListResourceTypesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The query conditions.
	Query []*string `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	// The resource type.
	//
	// For more information about the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListResourceTypesRequest) GetQuery() []*string {
	return s.Query
}

func (s *ListResourceTypesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypesRequest) SetAcceptLanguage(v string) *ListResourceTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesRequest) SetQuery(v []*string) *ListResourceTypesRequest {
	s.Query = v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
