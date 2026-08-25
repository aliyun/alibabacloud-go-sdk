// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCredentialResponseBody
	GetCode() *string
	SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody
	GetData() *UpdateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *UpdateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCredentialResponseBody
	GetSuccess() *bool
}

type UpdateCredentialResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCredentialResponseBody) GetData() *UpdateCredentialResponseBodyData {
	return s.Data
}

func (s *UpdateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCredentialResponseBody) SetCode(v string) *UpdateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCredentialResponseBody) SetHttpStatusCode(v int32) *UpdateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetMessage(v string) *UpdateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetRequestId(v string) *UpdateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetSuccess(v bool) *UpdateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialResponseBodyData struct {
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// cred-123456
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// example:
	//
	// {"apiKey":"****************"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// example:
	//
	// 线上环境调用模型服务使用的 API Key
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *UpdateCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *UpdateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCredentialResponseBodyData) SetCreatedAt(v string) *UpdateCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialId(v string) *UpdateCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialMetadata(v string) *UpdateCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialType(v string) *UpdateCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetDescription(v string) *UpdateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetName(v string) *UpdateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetRegionId(v string) *UpdateCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetUpdatedAt(v string) *UpdateCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetWorkspaceId(v string) *UpdateCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
