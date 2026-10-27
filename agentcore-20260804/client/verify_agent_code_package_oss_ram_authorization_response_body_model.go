// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgentCodePackageOssRamAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetCode() *string
	SetData(v *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetData() *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData
	SetHttpStatusCode(v int32) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyAgentCodePackageOssRamAuthorizationResponseBody
	GetSuccess() *bool
}

type VerifyAgentCodePackageOssRamAuthorizationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s VerifyAgentCodePackageOssRamAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgentCodePackageOssRamAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetData() *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData {
	return s.Data
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetCode(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetData(v *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetHttpStatusCode(v int32) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetMessage(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetRequestId(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) SetSuccess(v bool) *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyAgentCodePackageOssRamAuthorizationResponseBodyData struct {
	// The policy binding status.
	//
	// example:
	//
	// AUTHORIZED
	AuthorizationStatus *string `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
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

func (s VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) SetAuthorizationStatus(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData {
	s.AuthorizationStatus = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) SetOssBucketName(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData {
	s.OssBucketName = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) SetRoleName(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) SetWorkspaceId(v string) *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
