// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryNamespaceRequest
	GetAcceptLanguage() *string
	SetName(v string) *QueryNamespaceRequest
	GetName() *string
	SetRegion(v string) *QueryNamespaceRequest
	GetRegion() *string
}

type QueryNamespaceRequest struct {
	// The language of the response. Valid values: zh and en. Default value: zh. The value zh which indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNamespaceRequest) GoString() string {
	return s.String()
}

func (s *QueryNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *QueryNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryNamespaceRequest) SetAcceptLanguage(v string) *QueryNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryNamespaceRequest) SetName(v string) *QueryNamespaceRequest {
	s.Name = &v
	return s
}

func (s *QueryNamespaceRequest) SetRegion(v string) *QueryNamespaceRequest {
	s.Region = &v
	return s
}

func (s *QueryNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
