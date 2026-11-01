// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAccessApplicationResponseBody
	GetCode() *string
	SetData(v *GetAccessApplicationResponseBodyData) *GetAccessApplicationResponseBody
	GetData() *GetAccessApplicationResponseBodyData
	SetHttpStatusCode(v int32) *GetAccessApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAccessApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccessApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAccessApplicationResponseBody
	GetSuccess() *bool
}

type GetAccessApplicationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The access application details.
	Data *GetAccessApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message about the request processing result.
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
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccessApplicationResponseBody) GetData() *GetAccessApplicationResponseBodyData {
	return s.Data
}

func (s *GetAccessApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAccessApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAccessApplicationResponseBody) SetCode(v string) *GetAccessApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessApplicationResponseBody) SetData(v *GetAccessApplicationResponseBodyData) *GetAccessApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetAccessApplicationResponseBody) SetHttpStatusCode(v int32) *GetAccessApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAccessApplicationResponseBody) SetMessage(v string) *GetAccessApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccessApplicationResponseBody) SetRequestId(v string) *GetAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessApplicationResponseBody) SetSuccess(v bool) *GetAccessApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *GetAccessApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessApplicationResponseBodyData struct {
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
	// RAM
	UserSource *string `json:"userSource,omitempty" xml:"userSource,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetAccessApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccessApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccessApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetAccessApplicationResponseBodyData) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetAccessApplicationResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetAccessApplicationResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAccessApplicationResponseBodyData) GetMetadataUrl() *string {
	return s.MetadataUrl
}

func (s *GetAccessApplicationResponseBodyData) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *GetAccessApplicationResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAccessApplicationResponseBodyData) GetUserSource() *string {
	return s.UserSource
}

func (s *GetAccessApplicationResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAccessApplicationResponseBodyData) SetApplicationId(v string) *GetAccessApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetApplicationName(v string) *GetAccessApplicationResponseBodyData {
	s.ApplicationName = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetClientId(v string) *GetAccessApplicationResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetCreatedAt(v string) *GetAccessApplicationResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetMetadataUrl(v string) *GetAccessApplicationResponseBodyData {
	s.MetadataUrl = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetRedirectUris(v []*string) *GetAccessApplicationResponseBodyData {
	s.RedirectUris = v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetUpdatedAt(v string) *GetAccessApplicationResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetUserSource(v string) *GetAccessApplicationResponseBodyData {
	s.UserSource = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) SetWorkspaceId(v string) *GetAccessApplicationResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetAccessApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
