// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterConnectionTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListClusterConnectionTypesRequest
	GetAcceptLanguage() *string
}

type ListClusterConnectionTypesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListClusterConnectionTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConnectionTypesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListClusterConnectionTypesRequest) SetAcceptLanguage(v string) *ListClusterConnectionTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListClusterConnectionTypesRequest) Validate() error {
	return dara.Validate(s)
}
