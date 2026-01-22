// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedScopesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPredefinedScopes(v *ListPredefinedScopesResponseBodyPredefinedScopes) *ListPredefinedScopesResponseBody
	GetPredefinedScopes() *ListPredefinedScopesResponseBodyPredefinedScopes
	SetRequestId(v string) *ListPredefinedScopesResponseBody
	GetRequestId() *string
}

type ListPredefinedScopesResponseBody struct {
	// The information about application permissions.
	PredefinedScopes *ListPredefinedScopesResponseBodyPredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 955C096D-EC99-480B-AF37-3921109107D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPredefinedScopesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBody) GetPredefinedScopes() *ListPredefinedScopesResponseBodyPredefinedScopes {
	return s.PredefinedScopes
}

func (s *ListPredefinedScopesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPredefinedScopesResponseBody) SetPredefinedScopes(v *ListPredefinedScopesResponseBodyPredefinedScopes) *ListPredefinedScopesResponseBody {
	s.PredefinedScopes = v
	return s
}

func (s *ListPredefinedScopesResponseBody) SetRequestId(v string) *ListPredefinedScopesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPredefinedScopesResponseBody) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPredefinedScopesResponseBodyPredefinedScopes struct {
	PredefinedScope []*ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListPredefinedScopesResponseBodyPredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedScopesResponseBodyPredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopes) GetPredefinedScope() []*ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopes) SetPredefinedScope(v []*ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) *ListPredefinedScopesResponseBodyPredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopes) Validate() error {
	if s.PredefinedScope != nil {
		for _, item := range s.PredefinedScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope struct {
	// The description of the permission scope.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission scope.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) SetDescription(v string) *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) SetName(v string) *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
