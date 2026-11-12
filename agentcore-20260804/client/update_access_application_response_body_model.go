// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAccessApplicationResponseBody
	GetCode() *string
	SetData(v *UpdateAccessApplicationResponseBodyData) *UpdateAccessApplicationResponseBody
	GetData() *UpdateAccessApplicationResponseBodyData
	SetHttpStatusCode(v int32) *UpdateAccessApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAccessApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAccessApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAccessApplicationResponseBody
	GetSuccess() *bool
}

type UpdateAccessApplicationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the access application.
	Data *UpdateAccessApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message indicating the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAccessApplicationResponseBody) GetData() *UpdateAccessApplicationResponseBodyData {
	return s.Data
}

func (s *UpdateAccessApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAccessApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccessApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAccessApplicationResponseBody) SetCode(v string) *UpdateAccessApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccessApplicationResponseBody) SetData(v *UpdateAccessApplicationResponseBodyData) *UpdateAccessApplicationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAccessApplicationResponseBody) SetHttpStatusCode(v int32) *UpdateAccessApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAccessApplicationResponseBody) SetMessage(v string) *UpdateAccessApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccessApplicationResponseBody) SetRequestId(v string) *UpdateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessApplicationResponseBody) SetSuccess(v bool) *UpdateAccessApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAccessApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAccessApplicationResponseBodyData struct {
	// The access application ID.
	//
	// example:
	//
	// app-1234567890
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the access application.
	//
	// example:
	//
	// my-access-application
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The OAuth client ID.
	//
	// example:
	//
	// client-1234567890
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The creation time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The metadata URL of the identity provider.
	//
	// example:
	//
	// https://example.com/.well-known/openid-configuration
	MetadataUrl *string `json:"metadataUrl,omitempty" xml:"metadataUrl,omitempty"`
	// The list of redirect URIs.
	RedirectUris []*string `json:"redirectUris,omitempty" xml:"redirectUris,omitempty" type:"Repeated"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The user source.
	//
	// example:
	//
	// LOCAL
	UserSource *string `json:"userSource,omitempty" xml:"userSource,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateAccessApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateAccessApplicationResponseBodyData) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateAccessApplicationResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateAccessApplicationResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateAccessApplicationResponseBodyData) GetMetadataUrl() *string {
	return s.MetadataUrl
}

func (s *UpdateAccessApplicationResponseBodyData) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *UpdateAccessApplicationResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateAccessApplicationResponseBodyData) GetUserSource() *string {
	return s.UserSource
}

func (s *UpdateAccessApplicationResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAccessApplicationResponseBodyData) SetApplicationId(v string) *UpdateAccessApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetApplicationName(v string) *UpdateAccessApplicationResponseBodyData {
	s.ApplicationName = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetClientId(v string) *UpdateAccessApplicationResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetCreatedAt(v string) *UpdateAccessApplicationResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetMetadataUrl(v string) *UpdateAccessApplicationResponseBodyData {
	s.MetadataUrl = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetRedirectUris(v []*string) *UpdateAccessApplicationResponseBodyData {
	s.RedirectUris = v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetUpdatedAt(v string) *UpdateAccessApplicationResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetUserSource(v string) *UpdateAccessApplicationResponseBodyData {
	s.UserSource = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) SetWorkspaceId(v string) *UpdateAccessApplicationResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAccessApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
