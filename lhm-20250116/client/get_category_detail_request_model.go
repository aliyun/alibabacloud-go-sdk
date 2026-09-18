// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoryDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetCategoryDetailRequest
	GetCategory() *string
	SetSource(v string) *GetCategoryDetailRequest
	GetSource() *string
	SetTarget(v string) *GetCategoryDetailRequest
	GetTarget() *string
}

type GetCategoryDetailRequest struct {
	// The list of category paths.
	//
	// example:
	//
	// /lhm/
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The source dialect.
	//
	// example:
	//
	// sparksql
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target dialect.
	//
	// example:
	//
	// hologres
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s GetCategoryDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCategoryDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCategoryDetailRequest) GetCategory() *string {
	return s.Category
}

func (s *GetCategoryDetailRequest) GetSource() *string {
	return s.Source
}

func (s *GetCategoryDetailRequest) GetTarget() *string {
	return s.Target
}

func (s *GetCategoryDetailRequest) SetCategory(v string) *GetCategoryDetailRequest {
	s.Category = &v
	return s
}

func (s *GetCategoryDetailRequest) SetSource(v string) *GetCategoryDetailRequest {
	s.Source = &v
	return s
}

func (s *GetCategoryDetailRequest) SetTarget(v string) *GetCategoryDetailRequest {
	s.Target = &v
	return s
}

func (s *GetCategoryDetailRequest) Validate() error {
	return dara.Validate(s)
}
