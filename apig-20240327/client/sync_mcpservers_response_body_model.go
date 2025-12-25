// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMCPServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncMCPServersResponseBody
	GetCode() *string
	SetData(v *SyncMCPServersResponseBodyData) *SyncMCPServersResponseBody
	GetData() *SyncMCPServersResponseBodyData
	SetMessage(v string) *SyncMCPServersResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncMCPServersResponseBody
	GetRequestId() *string
}

type SyncMCPServersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *SyncMCPServersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SyncMCPServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMCPServersResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncMCPServersResponseBody) GetData() *SyncMCPServersResponseBodyData {
	return s.Data
}

func (s *SyncMCPServersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncMCPServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncMCPServersResponseBody) SetCode(v string) *SyncMCPServersResponseBody {
	s.Code = &v
	return s
}

func (s *SyncMCPServersResponseBody) SetData(v *SyncMCPServersResponseBodyData) *SyncMCPServersResponseBody {
	s.Data = v
	return s
}

func (s *SyncMCPServersResponseBody) SetMessage(v string) *SyncMCPServersResponseBody {
	s.Message = &v
	return s
}

func (s *SyncMCPServersResponseBody) SetRequestId(v string) *SyncMCPServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMCPServersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncMCPServersResponseBodyData struct {
	FailedMcpServers  []*SyncMCPServersResponseBodyDataFailedMcpServers  `json:"failedMcpServers,omitempty" xml:"failedMcpServers,omitempty" type:"Repeated"`
	SucceedMcpServers []*SyncMCPServersResponseBodyDataSucceedMcpServers `json:"succeedMcpServers,omitempty" xml:"succeedMcpServers,omitempty" type:"Repeated"`
}

func (s SyncMCPServersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncMCPServersResponseBodyData) GetFailedMcpServers() []*SyncMCPServersResponseBodyDataFailedMcpServers {
	return s.FailedMcpServers
}

func (s *SyncMCPServersResponseBodyData) GetSucceedMcpServers() []*SyncMCPServersResponseBodyDataSucceedMcpServers {
	return s.SucceedMcpServers
}

func (s *SyncMCPServersResponseBodyData) SetFailedMcpServers(v []*SyncMCPServersResponseBodyDataFailedMcpServers) *SyncMCPServersResponseBodyData {
	s.FailedMcpServers = v
	return s
}

func (s *SyncMCPServersResponseBodyData) SetSucceedMcpServers(v []*SyncMCPServersResponseBodyDataSucceedMcpServers) *SyncMCPServersResponseBodyData {
	s.SucceedMcpServers = v
	return s
}

func (s *SyncMCPServersResponseBodyData) Validate() error {
	if s.FailedMcpServers != nil {
		for _, item := range s.FailedMcpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SucceedMcpServers != nil {
		for _, item := range s.SucceedMcpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncMCPServersResponseBodyDataFailedMcpServers struct {
	// example:
	//
	// mcp-fail
	McpServerName *string   `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	Protocols     []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
}

func (s SyncMCPServersResponseBodyDataFailedMcpServers) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersResponseBodyDataFailedMcpServers) GoString() string {
	return s.String()
}

func (s *SyncMCPServersResponseBodyDataFailedMcpServers) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *SyncMCPServersResponseBodyDataFailedMcpServers) GetProtocols() []*string {
	return s.Protocols
}

func (s *SyncMCPServersResponseBodyDataFailedMcpServers) SetMcpServerName(v string) *SyncMCPServersResponseBodyDataFailedMcpServers {
	s.McpServerName = &v
	return s
}

func (s *SyncMCPServersResponseBodyDataFailedMcpServers) SetProtocols(v []*string) *SyncMCPServersResponseBodyDataFailedMcpServers {
	s.Protocols = v
	return s
}

func (s *SyncMCPServersResponseBodyDataFailedMcpServers) Validate() error {
	return dara.Validate(s)
}

type SyncMCPServersResponseBodyDataSucceedMcpServers struct {
	// example:
	//
	// mcp-success
	McpServerName *string   `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	Protocols     []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
}

func (s SyncMCPServersResponseBodyDataSucceedMcpServers) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersResponseBodyDataSucceedMcpServers) GoString() string {
	return s.String()
}

func (s *SyncMCPServersResponseBodyDataSucceedMcpServers) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *SyncMCPServersResponseBodyDataSucceedMcpServers) GetProtocols() []*string {
	return s.Protocols
}

func (s *SyncMCPServersResponseBodyDataSucceedMcpServers) SetMcpServerName(v string) *SyncMCPServersResponseBodyDataSucceedMcpServers {
	s.McpServerName = &v
	return s
}

func (s *SyncMCPServersResponseBodyDataSucceedMcpServers) SetProtocols(v []*string) *SyncMCPServersResponseBodyDataSucceedMcpServers {
	s.Protocols = v
	return s
}

func (s *SyncMCPServersResponseBodyDataSucceedMcpServers) Validate() error {
	return dara.Validate(s)
}
