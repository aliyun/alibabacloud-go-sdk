// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeEsRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *InvokeEsRequestRequest
	GetBody() *string
	SetCredentialId(v string) *InvokeEsRequestRequest
	GetCredentialId() *string
	SetMethod(v string) *InvokeEsRequestRequest
	GetMethod() *string
	SetPath(v string) *InvokeEsRequestRequest
	GetPath() *string
	SetSystem(v bool) *InvokeEsRequestRequest
	GetSystem() *bool
}

type InvokeEsRequestRequest struct {
	// The request body passed through to ES. Set this parameter based on the requirements of the target ES API. This parameter is not required for calls such as GET that do not have a request body.
	//
	// example:
	//
	// {"query":{"match_all":{}}}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the credential to use. If this parameter is not specified, the default credential of the instance is used.
	//
	// example:
	//
	// cred-7k2mq9xr4vbn
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The HTTP method used to access ES. Default value: GET.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The ES path to access. This parameter is required. The leading / can be omitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// _cat/indices?format=json
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to use the Alibaba Cloud ES system credential. Default value: false.
	//
	// example:
	//
	// false
	System *bool `json:"system,omitempty" xml:"system,omitempty"`
}

func (s InvokeEsRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeEsRequestRequest) GoString() string {
	return s.String()
}

func (s *InvokeEsRequestRequest) GetBody() *string {
	return s.Body
}

func (s *InvokeEsRequestRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *InvokeEsRequestRequest) GetMethod() *string {
	return s.Method
}

func (s *InvokeEsRequestRequest) GetPath() *string {
	return s.Path
}

func (s *InvokeEsRequestRequest) GetSystem() *bool {
	return s.System
}

func (s *InvokeEsRequestRequest) SetBody(v string) *InvokeEsRequestRequest {
	s.Body = &v
	return s
}

func (s *InvokeEsRequestRequest) SetCredentialId(v string) *InvokeEsRequestRequest {
	s.CredentialId = &v
	return s
}

func (s *InvokeEsRequestRequest) SetMethod(v string) *InvokeEsRequestRequest {
	s.Method = &v
	return s
}

func (s *InvokeEsRequestRequest) SetPath(v string) *InvokeEsRequestRequest {
	s.Path = &v
	return s
}

func (s *InvokeEsRequestRequest) SetSystem(v bool) *InvokeEsRequestRequest {
	s.System = &v
	return s
}

func (s *InvokeEsRequestRequest) Validate() error {
	return dara.Validate(s)
}
