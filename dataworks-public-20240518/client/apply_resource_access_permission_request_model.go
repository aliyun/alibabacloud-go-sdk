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
	// The list of resource permission application contents.
	//
	// This parameter is required.
	ApplyContents []*ApplyResourceAccessPermissionRequestApplyContents `json:"ApplyContents,omitempty" xml:"ApplyContents,omitempty" type:"Repeated"`
	// The idempotency parameter. Used to prevent duplicate operations caused by multiple calls.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The reason for the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 业务发展需要
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
	// The list of permissions to apply for.
	//
	// **Note**: Different resource levels support different permission types. They are uniformly constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).isValidLeaf, accessTypeRestrictions, and authMethodAccessTypes.
	//
	// Appendix: [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// The authorization method. Currently, only SEVERLESS_STARROCKS supports specifying the authorization method: ranger or starrocksManager.
	//
	// **Note**: Different resources support different authorization methods, which are uniformly constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).authMethods.
	//
	// Appendix: [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// ranger
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// The permission expiration time, in milliseconds timestamp.
	//
	// example:
	//
	// 1785835708000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The grantee description.
	//
	// **Note**: The supported grantee types are constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).authPrincipal.
	//
	// Appendix: [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	Grantee *ApplyResourceAccessPermissionRequestApplyContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// The resource description.
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
	// The grantee ID. The ID has different semantics depending on the grantee type:
	//
	// - RamUser: Dataworks UserId
	//
	// - RamRole: Dataworks UserId prefixed with "ROLE_"
	//
	// - DlfRole: DlfNext role name
	//
	// This parameter is required.
	//
	// example:
	//
	// ROLE_32237475848545
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The grantee type. Valid values:
	//
	// - RamRole
	//
	// - RamUser
	//
	// - DlfRole
	//
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
	// The resource type.
	//
	// **Note**: The resource types supported for application are constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).name.
	//
	// Appendix: [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The resource parsing version, which is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).version.
	//
	// [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// The resource metadata declaration.
	//
	// **Note**: The metadata is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).resources. A valid resource declaration must include full-path metadata declarations from level 0 to validLeaf.
	//
	// Appendix: [ResourceSchema documentation for international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	MetaData map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
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
