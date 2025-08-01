// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNamespaceShrinkRequest
	GetAcceptLanguage() *string
	SetDescribe(v string) *CreateNamespaceShrinkRequest
	GetDescribe() *string
	SetName(v string) *CreateNamespaceShrinkRequest
	GetName() *string
	SetTagShrink(v string) *CreateNamespaceShrinkRequest
	GetTagShrink() *string
}

type CreateNamespaceShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	Describe       *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// example:
	//
	// myNamespace
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateNamespaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNamespaceShrinkRequest) GetDescribe() *string {
	return s.Describe
}

func (s *CreateNamespaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateNamespaceShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateNamespaceShrinkRequest) SetAcceptLanguage(v string) *CreateNamespaceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetDescribe(v string) *CreateNamespaceShrinkRequest {
	s.Describe = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetName(v string) *CreateNamespaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetTagShrink(v string) *CreateNamespaceShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
