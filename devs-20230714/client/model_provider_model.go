// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProvider interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *ModelProvider
	GetCreatedTime() *string
	SetDescription(v string) *ModelProvider
	GetDescription() *string
	SetKind(v string) *ModelProvider
	GetKind() *string
	SetLabels(v map[string]*string) *ModelProvider
	GetLabels() map[string]*string
	SetName(v string) *ModelProvider
	GetName() *string
	SetUid(v string) *ModelProvider
	GetUid() *string
}

type ModelProvider struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Toolset
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-toolset
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ModelProvider) String() string {
	return dara.Prettify(s)
}

func (s ModelProvider) GoString() string {
	return s.String()
}

func (s *ModelProvider) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ModelProvider) GetDescription() *string {
	return s.Description
}

func (s *ModelProvider) GetKind() *string {
	return s.Kind
}

func (s *ModelProvider) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ModelProvider) GetName() *string {
	return s.Name
}

func (s *ModelProvider) GetUid() *string {
	return s.Uid
}

func (s *ModelProvider) SetCreatedTime(v string) *ModelProvider {
	s.CreatedTime = &v
	return s
}

func (s *ModelProvider) SetDescription(v string) *ModelProvider {
	s.Description = &v
	return s
}

func (s *ModelProvider) SetKind(v string) *ModelProvider {
	s.Kind = &v
	return s
}

func (s *ModelProvider) SetLabels(v map[string]*string) *ModelProvider {
	s.Labels = v
	return s
}

func (s *ModelProvider) SetName(v string) *ModelProvider {
	s.Name = &v
	return s
}

func (s *ModelProvider) SetUid(v string) *ModelProvider {
	s.Uid = &v
	return s
}

func (s *ModelProvider) Validate() error {
	return dara.Validate(s)
}
