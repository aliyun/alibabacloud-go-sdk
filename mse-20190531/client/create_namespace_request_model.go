// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNamespaceRequest
	GetAcceptLanguage() *string
	SetDescribe(v string) *CreateNamespaceRequest
	GetDescribe() *string
	SetName(v string) *CreateNamespaceRequest
	GetName() *string
	SetTag(v []*CreateNamespaceRequestTag) *CreateNamespaceRequest
	GetTag() []*CreateNamespaceRequestTag
}

type CreateNamespaceRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	Describe       *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// example:
	//
	// myNamespace
	Name *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag  []*CreateNamespaceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNamespaceRequest) GetDescribe() *string {
	return s.Describe
}

func (s *CreateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNamespaceRequest) GetTag() []*CreateNamespaceRequestTag {
	return s.Tag
}

func (s *CreateNamespaceRequest) SetAcceptLanguage(v string) *CreateNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNamespaceRequest) SetDescribe(v string) *CreateNamespaceRequest {
	s.Describe = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetTag(v []*CreateNamespaceRequestTag) *CreateNamespaceRequest {
	s.Tag = v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNamespaceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNamespaceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNamespaceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNamespaceRequestTag) SetKey(v string) *CreateNamespaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNamespaceRequestTag) SetValue(v string) *CreateNamespaceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNamespaceRequestTag) Validate() error {
	return dara.Validate(s)
}
