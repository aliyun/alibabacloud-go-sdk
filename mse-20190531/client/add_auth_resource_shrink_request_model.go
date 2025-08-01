// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddAuthResourceShrinkRequest
	GetAcceptLanguage() *string
	SetAuthId(v int64) *AddAuthResourceShrinkRequest
	GetAuthId() *int64
	SetAuthResourceHeaderListShrink(v string) *AddAuthResourceShrinkRequest
	GetAuthResourceHeaderListShrink() *string
	SetDomainId(v int64) *AddAuthResourceShrinkRequest
	GetDomainId() *int64
	SetGatewayUniqueId(v string) *AddAuthResourceShrinkRequest
	GetGatewayUniqueId() *string
	SetIgnoreCase(v bool) *AddAuthResourceShrinkRequest
	GetIgnoreCase() *bool
	SetMatchType(v string) *AddAuthResourceShrinkRequest
	GetMatchType() *string
	SetPath(v string) *AddAuthResourceShrinkRequest
	GetPath() *string
}

type AddAuthResourceShrinkRequest struct {
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
	// The ID of the authorization record.
	//
	// example:
	//
	// 13
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The authentication resource headers.
	AuthResourceHeaderListShrink *string `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 128
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether the matching is not case-sensitive. Default value: true.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The matching type. Valid values:
	//
	// 	- EQUAL
	//
	// 	- PRE
	//
	// 	- ERGULAR
	//
	// example:
	//
	// exact
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The path.
	//
	// example:
	//
	// /abc
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AddAuthResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuthResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAuthResourceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddAuthResourceShrinkRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *AddAuthResourceShrinkRequest) GetAuthResourceHeaderListShrink() *string {
	return s.AuthResourceHeaderListShrink
}

func (s *AddAuthResourceShrinkRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddAuthResourceShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddAuthResourceShrinkRequest) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *AddAuthResourceShrinkRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *AddAuthResourceShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *AddAuthResourceShrinkRequest) SetAcceptLanguage(v string) *AddAuthResourceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetAuthId(v int64) *AddAuthResourceShrinkRequest {
	s.AuthId = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetAuthResourceHeaderListShrink(v string) *AddAuthResourceShrinkRequest {
	s.AuthResourceHeaderListShrink = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetDomainId(v int64) *AddAuthResourceShrinkRequest {
	s.DomainId = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetGatewayUniqueId(v string) *AddAuthResourceShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetIgnoreCase(v bool) *AddAuthResourceShrinkRequest {
	s.IgnoreCase = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetMatchType(v string) *AddAuthResourceShrinkRequest {
	s.MatchType = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) SetPath(v string) *AddAuthResourceShrinkRequest {
	s.Path = &v
	return s
}

func (s *AddAuthResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
