// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentBootstrapOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetExternalAgentBootstrapOptionsResponseBody
	GetCode() *string
	SetData(v *GetExternalAgentBootstrapOptionsResponseBodyData) *GetExternalAgentBootstrapOptionsResponseBody
	GetData() *GetExternalAgentBootstrapOptionsResponseBodyData
	SetHttpStatusCode(v int32) *GetExternalAgentBootstrapOptionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetExternalAgentBootstrapOptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExternalAgentBootstrapOptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExternalAgentBootstrapOptionsResponseBody
	GetSuccess() *bool
}

type GetExternalAgentBootstrapOptionsResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The available network access information for the external agent.
	Data *GetExternalAgentBootstrapOptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request.
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

func (s GetExternalAgentBootstrapOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentBootstrapOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetData() *GetExternalAgentBootstrapOptionsResponseBodyData {
	return s.Data
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetCode(v string) *GetExternalAgentBootstrapOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetData(v *GetExternalAgentBootstrapOptionsResponseBodyData) *GetExternalAgentBootstrapOptionsResponseBody {
	s.Data = v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetHttpStatusCode(v int32) *GetExternalAgentBootstrapOptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetMessage(v string) *GetExternalAgentBootstrapOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetRequestId(v string) *GetExternalAgentBootstrapOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) SetSuccess(v bool) *GetExternalAgentBootstrapOptionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentBootstrapOptionsResponseBodyData struct {
	// The external agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The list of available network access options.
	NetworkOptions []*GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions `json:"networkOptions,omitempty" xml:"networkOptions,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetExternalAgentBootstrapOptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentBootstrapOptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) GetNetworkOptions() []*GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions {
	return s.NetworkOptions
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) SetAgentId(v string) *GetExternalAgentBootstrapOptionsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) SetNetworkOptions(v []*GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) *GetExternalAgentBootstrapOptionsResponseBodyData {
	s.NetworkOptions = v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) SetWorkspaceId(v string) *GetExternalAgentBootstrapOptionsResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyData) Validate() error {
	if s.NetworkOptions != nil {
		for _, item := range s.NetworkOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions struct {
	// Indicates whether the network access type is available.
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The network type. Valid values:
	//
	// - INTRANET: internal network.
	//
	// - INTERNET: public network.
	//
	// example:
	//
	// INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) GoString() string {
	return s.String()
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) GetAvailable() *bool {
	return s.Available
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) SetAvailable(v bool) *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions {
	s.Available = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) SetNetworkType(v string) *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions {
	s.NetworkType = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponseBodyDataNetworkOptions) Validate() error {
	return dara.Validate(s)
}
