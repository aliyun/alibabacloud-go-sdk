// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetCode() *string
	SetData(v *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetData() *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
	GetSuccess() *bool
}

type GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetData() *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetCode(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetData(v *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetHttpStatusCode(v int32) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetMessage(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetRequestId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) SetSuccess(v bool) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData struct {
	// The RAM authorization URL. After opening this URL and completing the authorization, call the verification operation.
	//
	// example:
	//
	// https://ram.console.aliyun.com/authorize?request=%7B%22referrer%22%3A%22AgentCore%22%2C%22payloads%22%3A%5B%7B%22missionId%22%3A%22AgentCore.CustomRoleForOSSObjectRestore%22%2C%22roleName%22%3A%22AgentCoreWorkspaceRoleExample%22%2C%22rolePolicies%22%3A%5B%7B%22policyName%22%3A%22AgentCoreNasMountExample%22%2C%22policyType%22%3A%22Custom%22%2C%22templateId%22%3A%22NasFileSystemClientMount%22%2C%22templateValue%22%3A%7B%22nasFileSystem%22%3A%22acs%3Anas%3Acn-hangzhou%3A1234567890123456%3Afilesystem%2F0123456789%22%2C%22nasAccessPoint%22%3A%22acs%3Anas%3Acn-hangzhou%3A1234567890123456%3Aaccesspoint%2Fap-0123456789abcdef0%22%7D%7D%5D%7D%5D%7D
	AuthorizeUrl *string `json:"authorizeUrl,omitempty" xml:"authorizeUrl,omitempty"`
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) GetAuthorizeUrl() *string {
	return s.AuthorizeUrl
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) SetAuthorizeUrl(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData {
	s.AuthorizeUrl = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
