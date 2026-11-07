// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNamespaceRequest
	GetAcceptLanguage() *string
	SetName(v string) *DeleteNamespaceRequest
	GetName() *string
	SetRegion(v string) *DeleteNamespaceRequest
	GetRegion() *string
}

type DeleteNamespaceRequest struct {
	// The language in which the response is displayed. Valid values: zh (default): Chinese, en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the MSE namespace to be deleted.
	//
	// example:
	//
	// prod
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region of the operation.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *DeleteNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteNamespaceRequest) SetAcceptLanguage(v string) *DeleteNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNamespaceRequest) SetName(v string) *DeleteNamespaceRequest {
	s.Name = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegion(v string) *DeleteNamespaceRequest {
	s.Region = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
