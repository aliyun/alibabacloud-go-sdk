// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourceAccessPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyContents(v []*ApplyResourceAccessPermissionRequestApplyContents) *ApplyResourceAccessPermissionRequest
	GetApplyContents() []*ApplyResourceAccessPermissionRequestApplyContents
	SetClientToken(v string) *ApplyResourceAccessPermissionRequest
	GetClientToken() *string
	SetReason(v string) *ApplyResourceAccessPermissionRequest
	GetReason() *string
}

type ApplyResourceAccessPermissionRequest struct {
	// This parameter is required.
	ApplyContents []*ApplyResourceAccessPermissionRequestApplyContents `json:"ApplyContents,omitempty" xml:"ApplyContents,omitempty" type:"Repeated"`
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ApplyResourceAccessPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionRequest) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionRequest) GetApplyContents() []*ApplyResourceAccessPermissionRequestApplyContents {
	return s.ApplyContents
}

func (s *ApplyResourceAccessPermissionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyResourceAccessPermissionRequest) GetReason() *string {
	return s.Reason
}

func (s *ApplyResourceAccessPermissionRequest) SetApplyContents(v []*ApplyResourceAccessPermissionRequestApplyContents) *ApplyResourceAccessPermissionRequest {
	s.ApplyContents = v
	return s
}

func (s *ApplyResourceAccessPermissionRequest) SetClientToken(v string) *ApplyResourceAccessPermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequest) SetReason(v string) *ApplyResourceAccessPermissionRequest {
	s.Reason = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequest) Validate() error {
	if s.ApplyContents != nil {
		for _, item := range s.ApplyContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyResourceAccessPermissionRequestApplyContents struct {
	// This parameter is required.
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// example:
	//
	// ranger
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// example:
	//
	// 1785835708000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// This parameter is required.
	Grantee  *ApplyResourceAccessPermissionRequestApplyContentsGrantee  `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	Resource *ApplyResourceAccessPermissionRequestApplyContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s ApplyResourceAccessPermissionRequestApplyContents) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionRequestApplyContents) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) GetGrantee() *ApplyResourceAccessPermissionRequestApplyContentsGrantee {
	return s.Grantee
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) GetResource() *ApplyResourceAccessPermissionRequestApplyContentsResource {
	return s.Resource
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) SetAccessTypes(v []*string) *ApplyResourceAccessPermissionRequestApplyContents {
	s.AccessTypes = v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) SetAuthMethod(v string) *ApplyResourceAccessPermissionRequestApplyContents {
	s.AuthMethod = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) SetExpirationTime(v int64) *ApplyResourceAccessPermissionRequestApplyContents {
	s.ExpirationTime = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) SetGrantee(v *ApplyResourceAccessPermissionRequestApplyContentsGrantee) *ApplyResourceAccessPermissionRequestApplyContents {
	s.Grantee = v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) SetResource(v *ApplyResourceAccessPermissionRequestApplyContentsResource) *ApplyResourceAccessPermissionRequestApplyContents {
	s.Resource = v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContents) Validate() error {
	if s.Grantee != nil {
		if err := s.Grantee.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyResourceAccessPermissionRequestApplyContentsGrantee struct {
	// This parameter is required.
	//
	// example:
	//
	// ROLE_32237475848545
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ApplyResourceAccessPermissionRequestApplyContentsGrantee) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionRequestApplyContentsGrantee) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsGrantee) SetPrincipalId(v string) *ApplyResourceAccessPermissionRequestApplyContentsGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsGrantee) SetPrincipalType(v string) *ApplyResourceAccessPermissionRequestApplyContentsGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsGrantee) Validate() error {
	return dara.Validate(s)
}

type ApplyResourceAccessPermissionRequestApplyContentsResource struct {
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// v1.0.0
	DefVersion *string                `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	MetaData   map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ApplyResourceAccessPermissionRequestApplyContentsResource) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionRequestApplyContentsResource) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) SetDefSchema(v string) *ApplyResourceAccessPermissionRequestApplyContentsResource {
	s.DefSchema = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) SetDefVersion(v string) *ApplyResourceAccessPermissionRequestApplyContentsResource {
	s.DefVersion = &v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) SetMetaData(v map[string]interface{}) *ApplyResourceAccessPermissionRequestApplyContentsResource {
	s.MetaData = v
	return s
}

func (s *ApplyResourceAccessPermissionRequestApplyContentsResource) Validate() error {
	return dara.Validate(s)
}
