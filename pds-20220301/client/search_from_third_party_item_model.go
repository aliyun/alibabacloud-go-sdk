// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFromThirdPartyItem interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v string) *SearchFromThirdPartyItem
	GetAuthenticationType() *string
	SetExtra(v string) *SearchFromThirdPartyItem
	GetExtra() *string
	SetIdentity(v string) *SearchFromThirdPartyItem
	GetIdentity() *string
	SetOthers(v map[string]interface{}) *SearchFromThirdPartyItem
	GetOthers() map[string]interface{}
}

type SearchFromThirdPartyItem struct {
	AuthenticationType *string                `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	Extra              *string                `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity           *string                `json:"identity,omitempty" xml:"identity,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
}

func (s SearchFromThirdPartyItem) String() string {
	return dara.Prettify(s)
}

func (s SearchFromThirdPartyItem) GoString() string {
	return s.String()
}

func (s *SearchFromThirdPartyItem) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *SearchFromThirdPartyItem) GetExtra() *string {
	return s.Extra
}

func (s *SearchFromThirdPartyItem) GetIdentity() *string {
	return s.Identity
}

func (s *SearchFromThirdPartyItem) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *SearchFromThirdPartyItem) SetAuthenticationType(v string) *SearchFromThirdPartyItem {
	s.AuthenticationType = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetExtra(v string) *SearchFromThirdPartyItem {
	s.Extra = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetIdentity(v string) *SearchFromThirdPartyItem {
	s.Identity = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetOthers(v map[string]interface{}) *SearchFromThirdPartyItem {
	s.Others = v
	return s
}

func (s *SearchFromThirdPartyItem) Validate() error {
	return dara.Validate(s)
}
