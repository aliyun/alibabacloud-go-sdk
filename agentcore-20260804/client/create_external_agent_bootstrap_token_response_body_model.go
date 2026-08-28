// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentBootstrapTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateExternalAgentBootstrapTokenResponseBody
	GetCode() *string
	SetData(v *CreateExternalAgentBootstrapTokenResponseBodyData) *CreateExternalAgentBootstrapTokenResponseBody
	GetData() *CreateExternalAgentBootstrapTokenResponseBodyData
	SetHttpStatusCode(v int32) *CreateExternalAgentBootstrapTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateExternalAgentBootstrapTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExternalAgentBootstrapTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExternalAgentBootstrapTokenResponseBody
	GetSuccess() *bool
}

type CreateExternalAgentBootstrapTokenResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The Bootstrap Token and CMS configuration required for connecting the external agent.
	Data *CreateExternalAgentBootstrapTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
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
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateExternalAgentBootstrapTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentBootstrapTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetData() *CreateExternalAgentBootstrapTokenResponseBodyData {
	return s.Data
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetCode(v string) *CreateExternalAgentBootstrapTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetData(v *CreateExternalAgentBootstrapTokenResponseBodyData) *CreateExternalAgentBootstrapTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetHttpStatusCode(v int32) *CreateExternalAgentBootstrapTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetMessage(v string) *CreateExternalAgentBootstrapTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetRequestId(v string) *CreateExternalAgentBootstrapTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) SetSuccess(v bool) *CreateExternalAgentBootstrapTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentBootstrapTokenResponseBodyData struct {
	// The external agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The Bootstrap Token used for connecting the external agent.
	//
	// example:
	//
	// bootstrap-token
	BootstrapToken *string `json:"bootstrapToken,omitempty" xml:"bootstrapToken,omitempty"`
	// The CMS configuration used for connecting the external agent.
	Cms *CreateExternalAgentBootstrapTokenResponseBodyDataCms `json:"cms,omitempty" xml:"cms,omitempty" type:"Struct"`
	// The network type for connection. Valid values:
	//
	// - INTERNET: public network
	//
	// - INTRANET: internal network
	//
	// example:
	//
	// INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The fingerprint of the Bootstrap Token.
	//
	// example:
	//
	// fingerprint
	TokenFingerprint *string `json:"tokenFingerprint,omitempty" xml:"tokenFingerprint,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateExternalAgentBootstrapTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentBootstrapTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetBootstrapToken() *string {
	return s.BootstrapToken
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetCms() *CreateExternalAgentBootstrapTokenResponseBodyDataCms {
	return s.Cms
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetTokenFingerprint() *string {
	return s.TokenFingerprint
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetAgentId(v string) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetBootstrapToken(v string) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.BootstrapToken = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetCms(v *CreateExternalAgentBootstrapTokenResponseBodyDataCms) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.Cms = v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetNetworkType(v string) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetTokenFingerprint(v string) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.TokenFingerprint = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) SetWorkspaceId(v string) *CreateExternalAgentBootstrapTokenResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyData) Validate() error {
	if s.Cms != nil {
		if err := s.Cms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentBootstrapTokenResponseBodyDataCms struct {
	// The CMS reporting endpoint.
	//
	// example:
	//
	// https://public.example.com/apm/trace/opentelemetry
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The license key used for CMS connection.
	//
	// example:
	//
	// license-key
	LicenseKey *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
	// The CMS workspace name.
	//
	// example:
	//
	// cms-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateExternalAgentBootstrapTokenResponseBodyDataCms) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentBootstrapTokenResponseBodyDataCms) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) SetEndpoint(v string) *CreateExternalAgentBootstrapTokenResponseBodyDataCms {
	s.Endpoint = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) SetLicenseKey(v string) *CreateExternalAgentBootstrapTokenResponseBodyDataCms {
	s.LicenseKey = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) SetWorkspace(v string) *CreateExternalAgentBootstrapTokenResponseBodyDataCms {
	s.Workspace = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponseBodyDataCms) Validate() error {
	return dara.Validate(s)
}
