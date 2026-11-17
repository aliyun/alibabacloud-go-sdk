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
	SetLoadAll(v bool) *ListRegionsRequest
	GetLoadAll() *bool
}

type ListRegionsRequest struct {
	// Specifies whether to filter regions based on assets for which auditing is enabled. Valid values: true: Auditing is enabled. false: Auditing is not enabled. If you do not specify this parameter, this condition is not used.
	Audited *bool `json:"Audited,omitempty" xml:"Audited,omitempty"`
	// Specifies whether to filter regions based on assets for which detection is enabled. Valid values: true: Detection is enabled. false: Detection is not enabled. If you do not specify this parameter, this condition is not used.
	Identified *bool `json:"Identified,omitempty" xml:"Identified,omitempty"`
	// The language of the response. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to retrieve all regions.
	LoadAll *bool `json:"LoadAll,omitempty" xml:"LoadAll,omitempty"`
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

func (s *ListRegionsRequest) GetLoadAll() *bool {
	return s.LoadAll
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

func (s *ListRegionsRequest) SetLoadAll(v bool) *ListRegionsRequest {
	s.LoadAll = &v
	return s
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
