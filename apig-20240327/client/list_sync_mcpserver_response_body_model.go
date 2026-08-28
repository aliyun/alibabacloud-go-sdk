// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyncMCPServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSyncMCPServerResponseBody
	GetCode() *string
	SetData(v *ListSyncMCPServerResponseBodyData) *ListSyncMCPServerResponseBody
	GetData() *ListSyncMCPServerResponseBodyData
	SetMessage(v string) *ListSyncMCPServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSyncMCPServerResponseBody
	GetRequestId() *string
}

type ListSyncMCPServerResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListSyncMCPServerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 211B1C7E-DD93-58D3-AA4B-9B392B63258C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSyncMCPServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSyncMCPServerResponseBody) GoString() string {
	return s.String()
}

func (s *ListSyncMCPServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSyncMCPServerResponseBody) GetData() *ListSyncMCPServerResponseBodyData {
	return s.Data
}

func (s *ListSyncMCPServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSyncMCPServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSyncMCPServerResponseBody) SetCode(v string) *ListSyncMCPServerResponseBody {
	s.Code = &v
	return s
}

func (s *ListSyncMCPServerResponseBody) SetData(v *ListSyncMCPServerResponseBodyData) *ListSyncMCPServerResponseBody {
	s.Data = v
	return s
}

func (s *ListSyncMCPServerResponseBody) SetMessage(v string) *ListSyncMCPServerResponseBody {
	s.Message = &v
	return s
}

func (s *ListSyncMCPServerResponseBody) SetRequestId(v string) *ListSyncMCPServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSyncMCPServerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSyncMCPServerResponseBodyData struct {
	DomainIds  []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	McpServers []*string `json:"mcpServers,omitempty" xml:"mcpServers,omitempty" type:"Repeated"`
}

func (s ListSyncMCPServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSyncMCPServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSyncMCPServerResponseBodyData) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *ListSyncMCPServerResponseBodyData) GetMcpServers() []*string {
	return s.McpServers
}

func (s *ListSyncMCPServerResponseBodyData) SetDomainIds(v []*string) *ListSyncMCPServerResponseBodyData {
	s.DomainIds = v
	return s
}

func (s *ListSyncMCPServerResponseBodyData) SetMcpServers(v []*string) *ListSyncMCPServerResponseBodyData {
	s.McpServers = v
	return s
}

func (s *ListSyncMCPServerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
