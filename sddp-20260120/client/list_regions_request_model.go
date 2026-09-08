// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudited(v bool) *ListRegionsRequest
	GetAudited() *bool
	SetIdentified(v bool) *ListRegionsRequest
	GetIdentified() *bool
	SetLang(v string) *ListRegionsRequest
	GetLang() *string
}

type ListRegionsRequest struct {
	Audited    *bool `json:"Audited,omitempty" xml:"Audited,omitempty"`
	Identified *bool `json:"Identified,omitempty" xml:"Identified,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) GetAudited() *bool {
	return s.Audited
}

func (s *ListRegionsRequest) GetIdentified() *bool {
	return s.Identified
}

func (s *ListRegionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListRegionsRequest) SetAudited(v bool) *ListRegionsRequest {
	s.Audited = &v
	return s
}

func (s *ListRegionsRequest) SetIdentified(v bool) *ListRegionsRequest {
	s.Identified = &v
	return s
}

func (s *ListRegionsRequest) SetLang(v string) *ListRegionsRequest {
	s.Lang = &v
	return s
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
