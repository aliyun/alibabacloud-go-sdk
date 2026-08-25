// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCredentialResponseBody
	GetCode() *string
	SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody
	GetData() *CreateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *CreateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCredentialResponseBody
	GetSuccess() *bool
}

type CreateCredentialResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCredentialResponseBody) GetData() *CreateCredentialResponseBodyData {
	return s.Data
}

func (s *CreateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCredentialResponseBody) SetCode(v string) *CreateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCredentialResponseBody) SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *CreateCredentialResponseBody) SetHttpStatusCode(v int32) *CreateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCredentialResponseBody) SetMessage(v string) *CreateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetSuccess(v bool) *CreateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialResponseBodyData struct {
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

func (s CreateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *CreateCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCredentialResponseBodyData) SetCreatedAt(v string) *CreateCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialId(v string) *CreateCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialMetadata(v string) *CreateCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialType(v string) *CreateCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetDescription(v string) *CreateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetName(v string) *CreateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetRegionId(v string) *CreateCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetUpdatedAt(v string) *CreateCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetWorkspaceId(v string) *CreateCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
