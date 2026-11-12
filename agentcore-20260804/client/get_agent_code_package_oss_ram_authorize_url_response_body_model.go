// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageOssRamAuthorizeUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetCode() *string
	SetData(v *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetData() *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
	GetSuccess() *bool
}

type GetAgentCodePackageOssRamAuthorizeUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetData() *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData {
	return s.Data
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetCode(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetData(v *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetHttpStatusCode(v int32) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetMessage(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetRequestId(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) SetSuccess(v bool) *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData struct {
	// The RAM authorization URL.
	//
	// example:
	//
	// https://ram.console.aliyun.com/authorize?request=%7B%22referrer%22%3A%22AgentCore%22%2C%22payloads%22%3A%5B%7B%22missionId%22%3A%22AgentCore.CustomRoleForOSSObjectRestore%22%2C%22roleName%22%3A%22AliyunAgentCoreWorkspace-ws-a1b2c3d4e5f6%22%2C%22rolePolicies%22%3A%5B%7B%22policyName%22%3A%22agentcore-oss-mount-policy-my-agent-bucket%22%2C%22policyType%22%3A%22Custom%22%2C%22templateId%22%3A%22OSSObjectRestoreCustom%22%2C%22templateValue%22%3A%7B%22ossBucket%22%3A%22acs%3Aoss%3A%2A%3A1234567890123456%3Amy-agent-bucket%22%2C%22ossBucketObject%22%3A%22acs%3Aoss%3A%2A%3A1234567890123456%3Amy-agent-bucket%2F%2A%22%7D%7D%5D%7D%5D%7D
	AuthorizeUrl *string `json:"authorizeUrl,omitempty" xml:"authorizeUrl,omitempty"`
	// The name of the authorized source OSS bucket.
	//
	// example:
	//
	// my-agent-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// The name of the workspace role selected by the backend.
	//
	// example:
	//
	// AliyunAgentCoreWorkspace-ws-a1b2c3d4e5f6
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-a1b2c3d4e5f6
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) GetAuthorizeUrl() *string {
	return s.AuthorizeUrl
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) SetAuthorizeUrl(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData {
	s.AuthorizeUrl = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) SetOssBucketName(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData {
	s.OssBucketName = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) SetRoleName(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) SetWorkspaceId(v string) *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
