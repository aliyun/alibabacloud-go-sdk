// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalMode(v string) *SetHttpDDoSAttackProtectionResponseBody
	GetGlobalMode() *string
	SetRequestId(v string) *SetHttpDDoSAttackProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetHttpDDoSAttackProtectionResponseBody
	GetSiteId() *int64
}

type SetHttpDDoSAttackProtectionResponseBody struct {
	// The level of HTTP DDoS attack protection.
	//
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionResponseBody) GetGlobalMode() *string {
	return s.GlobalMode
}

func (s *SetHttpDDoSAttackProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHttpDDoSAttackProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetGlobalMode(v string) *SetHttpDDoSAttackProtectionResponseBody {
	s.GlobalMode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetRequestId(v string) *SetHttpDDoSAttackProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetSiteId(v int64) *SetHttpDDoSAttackProtectionResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
