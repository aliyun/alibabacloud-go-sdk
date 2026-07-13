// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpResponseBody
	GetCode() *string
	SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody
	GetData() *GetMcpResponseBodyData
	SetMessage(v string) *GetMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpResponseBody
	GetSuccess() *bool
}

type GetMcpResponseBody struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMcpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpResponseBody) GetData() *GetMcpResponseBodyData {
	return s.Data
}

func (s *GetMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpResponseBody) SetCode(v string) *GetMcpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpResponseBody) SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpResponseBody) SetMessage(v string) *GetMcpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpResponseBody) SetRequestId(v string) *GetMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpResponseBody) SetSuccess(v bool) *GetMcpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyData struct {
	Addresses       []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	AuthConfig      *string   `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	AuthEnabled     *bool     `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	CreateType      *string   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DeployStatus    *string   `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	McpServerConfig *string   `json:"McpServerConfig,omitempty" xml:"McpServerConfig,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SwaggerConfig   *string   `json:"SwaggerConfig,omitempty" xml:"SwaggerConfig,omitempty"`
	Url             *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *GetMcpResponseBodyData) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *GetMcpResponseBodyData) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *GetMcpResponseBodyData) GetCreateType() *string {
	return s.CreateType
}

func (s *GetMcpResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMcpResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *GetMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpResponseBodyData) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *GetMcpResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetMcpResponseBodyData) SetAddresses(v []*string) *GetMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *GetMcpResponseBodyData) SetAuthConfig(v string) *GetMcpResponseBodyData {
	s.AuthConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetAuthEnabled(v bool) *GetMcpResponseBodyData {
	s.AuthEnabled = &v
	return s
}

func (s *GetMcpResponseBodyData) SetCreateType(v string) *GetMcpResponseBodyData {
	s.CreateType = &v
	return s
}

func (s *GetMcpResponseBodyData) SetDeployStatus(v string) *GetMcpResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *GetMcpResponseBodyData) SetDescription(v string) *GetMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMcpResponseBodyData) SetId(v string) *GetMcpResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMcpResponseBodyData) SetMcpServerConfig(v string) *GetMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetName(v string) *GetMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyData) SetProtocol(v string) *GetMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetMcpResponseBodyData) SetSwaggerConfig(v string) *GetMcpResponseBodyData {
	s.SwaggerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetUrl(v string) *GetMcpResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetMcpResponseBodyData) Validate() error {
	return dara.Validate(s)
}
