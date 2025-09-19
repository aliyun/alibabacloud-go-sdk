// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromEcs(v bool) *GetCanTrySasRequest
	GetFromEcs() *bool
	SetLang(v string) *GetCanTrySasRequest
	GetLang() *string
}

type GetCanTrySasRequest struct {
	// Specifies whether the request is redirected from the Elastic Compute Service (ECS) console. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FromEcs *bool `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetCanTrySasRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasRequest) GoString() string {
	return s.String()
}

func (s *GetCanTrySasRequest) GetFromEcs() *bool {
	return s.FromEcs
}

func (s *GetCanTrySasRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCanTrySasRequest) SetFromEcs(v bool) *GetCanTrySasRequest {
	s.FromEcs = &v
	return s
}

func (s *GetCanTrySasRequest) SetLang(v string) *GetCanTrySasRequest {
	s.Lang = &v
	return s
}

func (s *GetCanTrySasRequest) Validate() error {
	return dara.Validate(s)
}
