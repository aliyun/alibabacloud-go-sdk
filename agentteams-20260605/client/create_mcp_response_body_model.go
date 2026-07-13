// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpResponseBody
	GetCode() *string
	SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody
	GetData() *CreateMcpResponseBodyData
	SetMessage(v string) *CreateMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpResponseBody
	GetSuccess() *bool
}

type CreateMcpResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateMcpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpResponseBody) GetData() *CreateMcpResponseBodyData {
	return s.Data
}

func (s *CreateMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpResponseBody) SetCode(v string) *CreateMcpResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpResponseBody) SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpResponseBody) SetMessage(v string) *CreateMcpResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpResponseBody) SetRequestId(v string) *CreateMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpResponseBody) SetSuccess(v bool) *CreateMcpResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpResponseBodyData struct {
	Addresses       []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	CreateType      *string   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DeployStatus    *string   `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	McpServerConfig *string   `json:"McpServerConfig,omitempty" xml:"McpServerConfig,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpResponseBodyData) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateMcpResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *CreateMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *CreateMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpResponseBodyData) SetAddresses(v []*string) *CreateMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *CreateMcpResponseBodyData) SetCreateType(v string) *CreateMcpResponseBodyData {
	s.CreateType = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetDeployStatus(v string) *CreateMcpResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetDescription(v string) *CreateMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetId(v string) *CreateMcpResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetMcpServerConfig(v string) *CreateMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetName(v string) *CreateMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetProtocol(v string) *CreateMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateMcpResponseBodyData) Validate() error {
	return dara.Validate(s)
}
