// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAcrRamAuthorizeUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetCode() *string
	SetData(v *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetData() *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceAcrRamAuthorizeUrlResponseBody
	GetSuccess() *bool
}

type GetWorkspaceAcrRamAuthorizeUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 12345678-1234-1234-1234-123456789012
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetData() *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetCode(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetData(v *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetHttpStatusCode(v int32) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetMessage(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetRequestId(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) SetSuccess(v bool) *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceAcrRamAuthorizeUrlResponseBodyData struct {
	// The ACR Enterprise instance ID.
	//
	// example:
	//
	// cri-1234567890abcdef
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The RAM authorization URL used to create or append ACR permissions for the shared role on the target repository.
	//
	// example:
	//
	// https://ram.console.aliyun.com/authorize?request=%7B%22referrer%22%3A%22AgentCore%22%2C%22payloads%22%3A%5B%7B%22missionId%22%3A%22AgentCore.CustomRoleForOSSObjectRestore%22%2C%22roleName%22%3A%22AliyunAgentCoreWorkspace-ws-1234567890abcdef12345%22%2C%22rolePolicies%22%3A%5B%7B%22policyName%22%3A%22AgentCoreAcrPull-63373f3802583663d49465d8061e410a899d8f2936b8fc0cb9ec905b71f273c6%22%2C%22policyType%22%3A%22Custom%22%2C%22templateId%22%3A%22CrPullInstanceImage%22%2C%22templateValue%22%3A%7B%22crInstance%22%3A%22acs%3Acr%3Acn-hangzhou%3A1234567890123456%3Ainstance%2Fcri-1234567890abcdef%22%2C%22crRepository%22%3A%22acs%3Acr%3Acn-hangzhou%3A1234567890123456%3Arepository%2Fcri-1234567890abcdef%2Fexample%2Fagent-image%22%7D%7D%5D%7D%5D%7D
	AuthorizeUrl *string `json:"authorizeUrl,omitempty" xml:"authorizeUrl,omitempty"`
	// The shared role name selected by the backend. This value is not editable on the frontend.
	//
	// example:
	//
	// AliyunAgentCoreWorkspace-ws-1234567890abcdef12345
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// The source of the shared role. This value does not indicate that authorization is complete.
	//
	// example:
	//
	// WORKSPACE_SHARED
	RoleSource *string `json:"roleSource,omitempty" xml:"roleSource,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GetAuthorizeUrl() *string {
	return s.AuthorizeUrl
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GetRoleSource() *string {
	return s.RoleSource
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) SetAcrInstanceId(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) SetAuthorizeUrl(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	s.AuthorizeUrl = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) SetRoleName(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) SetRoleSource(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	s.RoleSource = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) SetWorkspaceId(v string) *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
