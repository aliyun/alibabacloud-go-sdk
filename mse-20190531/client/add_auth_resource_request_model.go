// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddAuthResourceRequest
	GetAcceptLanguage() *string
	SetAuthId(v int64) *AddAuthResourceRequest
	GetAuthId() *int64
	SetAuthResourceHeaderList(v []*AddAuthResourceRequestAuthResourceHeaderList) *AddAuthResourceRequest
	GetAuthResourceHeaderList() []*AddAuthResourceRequestAuthResourceHeaderList
	SetDomainId(v int64) *AddAuthResourceRequest
	GetDomainId() *int64
	SetGatewayUniqueId(v string) *AddAuthResourceRequest
	GetGatewayUniqueId() *string
	SetIgnoreCase(v bool) *AddAuthResourceRequest
	GetIgnoreCase() *bool
	SetMatchType(v string) *AddAuthResourceRequest
	GetMatchType() *string
	SetPath(v string) *AddAuthResourceRequest
	GetPath() *string
}

type AddAuthResourceRequest struct {
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
	AuthResourceHeaderList []*AddAuthResourceRequestAuthResourceHeaderList `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty" type:"Repeated"`
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

func (s AddAuthResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuthResourceRequest) GoString() string {
	return s.String()
}

func (s *AddAuthResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddAuthResourceRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *AddAuthResourceRequest) GetAuthResourceHeaderList() []*AddAuthResourceRequestAuthResourceHeaderList {
	return s.AuthResourceHeaderList
}

func (s *AddAuthResourceRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddAuthResourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddAuthResourceRequest) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *AddAuthResourceRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *AddAuthResourceRequest) GetPath() *string {
	return s.Path
}

func (s *AddAuthResourceRequest) SetAcceptLanguage(v string) *AddAuthResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddAuthResourceRequest) SetAuthId(v int64) *AddAuthResourceRequest {
	s.AuthId = &v
	return s
}

func (s *AddAuthResourceRequest) SetAuthResourceHeaderList(v []*AddAuthResourceRequestAuthResourceHeaderList) *AddAuthResourceRequest {
	s.AuthResourceHeaderList = v
	return s
}

func (s *AddAuthResourceRequest) SetDomainId(v int64) *AddAuthResourceRequest {
	s.DomainId = &v
	return s
}

func (s *AddAuthResourceRequest) SetGatewayUniqueId(v string) *AddAuthResourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddAuthResourceRequest) SetIgnoreCase(v bool) *AddAuthResourceRequest {
	s.IgnoreCase = &v
	return s
}

func (s *AddAuthResourceRequest) SetMatchType(v string) *AddAuthResourceRequest {
	s.MatchType = &v
	return s
}

func (s *AddAuthResourceRequest) SetPath(v string) *AddAuthResourceRequest {
	s.Path = &v
	return s
}

func (s *AddAuthResourceRequest) Validate() error {
	if s.AuthResourceHeaderList != nil {
		for _, item := range s.AuthResourceHeaderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddAuthResourceRequestAuthResourceHeaderList struct {
	// The parameter of the HTTP header.
	//
	// example:
	//
	// Access-Control-Allow-Origin
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	// The header matching mode.
	//
	// Valid values:
	//
	// 	- SUFFIX
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- EXIST
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PREFIX
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- EQUAL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- EXCLUDE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NOT_EQUAL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NOT_EXIST
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REGREX
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- INCLUDE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// EQUAL
	HeaderMethod *string `json:"HeaderMethod,omitempty" xml:"HeaderMethod,omitempty"`
	// The parameter value of the HTTP header.
	//
	// example:
	//
	// *
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s AddAuthResourceRequestAuthResourceHeaderList) String() string {
	return dara.Prettify(s)
}

func (s AddAuthResourceRequestAuthResourceHeaderList) GoString() string {
	return s.String()
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) GetHeaderMethod() *string {
	return s.HeaderMethod
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) SetHeaderKey(v string) *AddAuthResourceRequestAuthResourceHeaderList {
	s.HeaderKey = &v
	return s
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) SetHeaderMethod(v string) *AddAuthResourceRequestAuthResourceHeaderList {
	s.HeaderMethod = &v
	return s
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) SetHeaderValue(v string) *AddAuthResourceRequestAuthResourceHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *AddAuthResourceRequestAuthResourceHeaderList) Validate() error {
	return dara.Validate(s)
}
