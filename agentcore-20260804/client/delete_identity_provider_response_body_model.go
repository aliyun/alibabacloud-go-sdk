// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIdentityProviderResponseBody
	GetCode() *string
	SetData(v *DeleteIdentityProviderResponseBodyData) *DeleteIdentityProviderResponseBody
	GetData() *DeleteIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *DeleteIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIdentityProviderResponseBody
	GetSuccess() *bool
}

type DeleteIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding information of the external identity provider after the unbinding operation is accepted.
	Data *DeleteIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIdentityProviderResponseBody) GetData() *DeleteIdentityProviderResponseBodyData {
	return s.Data
}

func (s *DeleteIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIdentityProviderResponseBody) SetCode(v string) *DeleteIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetData(v *DeleteIdentityProviderResponseBodyData) *DeleteIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetHttpStatusCode(v int32) *DeleteIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetMessage(v string) *DeleteIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetRequestId(v string) *DeleteIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetSuccess(v bool) *DeleteIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteIdentityProviderResponseBodyData struct {
	// The type of the external identity provider. Valid values: DingTalk, Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The status. Valid values:
	//
	// - CONFIGURED: The configuration has been accepted and is waiting for the user pool to be provisioned.
	//
	// - SYNCING: Organization members are being synchronized.
	//
	// - SYNCED: Organization member synchronization is complete.
	//
	// - READY: The binding is in effect.
	//
	// - SYNC_FAILED: Organization member synchronization failed.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - UPDATE_FAILED: Configuration update failed.
	//
	// - DISCONNECTING: The unbinding is in progress.
	//
	// - DISCONNECT_FAILED: The unbinding failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *DeleteIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *DeleteIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetStatus(v string) *DeleteIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetWorkspaceId(v string) *DeleteIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
