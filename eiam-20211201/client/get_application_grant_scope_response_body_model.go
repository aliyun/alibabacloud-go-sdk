// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGrantScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGrantScope(v *GetApplicationGrantScopeResponseBodyApplicationGrantScope) *GetApplicationGrantScopeResponseBody
	GetApplicationGrantScope() *GetApplicationGrantScopeResponseBodyApplicationGrantScope
	SetRequestId(v string) *GetApplicationGrantScopeResponseBody
	GetRequestId() *string
}

type GetApplicationGrantScopeResponseBody struct {
	// The permissions of the Developer API feature.
	ApplicationGrantScope *GetApplicationGrantScopeResponseBodyApplicationGrantScope `json:"ApplicationGrantScope,omitempty" xml:"ApplicationGrantScope,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationGrantScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGrantScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponseBody) GetApplicationGrantScope() *GetApplicationGrantScopeResponseBodyApplicationGrantScope {
	return s.ApplicationGrantScope
}

func (s *GetApplicationGrantScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationGrantScopeResponseBody) SetApplicationGrantScope(v *GetApplicationGrantScopeResponseBodyApplicationGrantScope) *GetApplicationGrantScopeResponseBody {
	s.ApplicationGrantScope = v
	return s
}

func (s *GetApplicationGrantScopeResponseBody) SetRequestId(v string) *GetApplicationGrantScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationGrantScopeResponseBody) Validate() error {
	if s.ApplicationGrantScope != nil {
		if err := s.ApplicationGrantScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationGrantScopeResponseBodyApplicationGrantScope struct {
	// The permissions of the Developer API feature.
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
}

func (s GetApplicationGrantScopeResponseBodyApplicationGrantScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGrantScopeResponseBodyApplicationGrantScope) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponseBodyApplicationGrantScope) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *GetApplicationGrantScopeResponseBodyApplicationGrantScope) SetGrantScopes(v []*string) *GetApplicationGrantScopeResponseBodyApplicationGrantScope {
	s.GrantScopes = v
	return s
}

func (s *GetApplicationGrantScopeResponseBodyApplicationGrantScope) Validate() error {
	return dara.Validate(s)
}
