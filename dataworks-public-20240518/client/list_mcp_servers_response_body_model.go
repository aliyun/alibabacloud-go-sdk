// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListMcpServersResponseBodyPagingInfo) *ListMcpServersResponseBody
	GetPagingInfo() *ListMcpServersResponseBodyPagingInfo
	SetRequestId(v string) *ListMcpServersResponseBody
	GetRequestId() *string
}

type ListMcpServersResponseBody struct {
	PagingInfo *ListMcpServersResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMcpServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBody) GetPagingInfo() *ListMcpServersResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListMcpServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpServersResponseBody) SetPagingInfo(v *ListMcpServersResponseBodyPagingInfo) *ListMcpServersResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListMcpServersResponseBody) SetRequestId(v string) *ListMcpServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpServersResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpServersResponseBodyPagingInfo struct {
	MaxResults *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	McpServers []*ListMcpServersResponseBodyPagingInfoMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	NextToken  *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcpServersResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyPagingInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpServersResponseBodyPagingInfo) GetMcpServers() []*ListMcpServersResponseBodyPagingInfoMcpServers {
	return s.McpServers
}

func (s *ListMcpServersResponseBodyPagingInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpServersResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMcpServersResponseBodyPagingInfo) SetMaxResults(v int32) *ListMcpServersResponseBodyPagingInfo {
	s.MaxResults = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfo) SetMcpServers(v []*ListMcpServersResponseBodyPagingInfoMcpServers) *ListMcpServersResponseBodyPagingInfo {
	s.McpServers = v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfo) SetNextToken(v string) *ListMcpServersResponseBodyPagingInfo {
	s.NextToken = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfo) SetTotalCount(v int32) *ListMcpServersResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfo) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpServersResponseBodyPagingInfoMcpServers struct {
	Config    *ListMcpServersResponseBodyPagingInfoMcpServersConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	CreatorId *string                                               `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMcpServersResponseBodyPagingInfoMcpServers) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyPagingInfoMcpServers) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetConfig() *ListMcpServersResponseBodyPagingInfoMcpServersConfig {
	return s.Config
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetName() *string {
	return s.Name
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetConfig(v *ListMcpServersResponseBodyPagingInfoMcpServersConfig) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.Config = v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetCreatorId(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.CreatorId = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetGmtCreateTime(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetGmtModifiedTime(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetModifierId(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.ModifierId = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetName(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.Name = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) SetVisibility(v string) *ListMcpServersResponseBodyPagingInfoMcpServers {
	s.Visibility = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServers) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpServersResponseBodyPagingInfoMcpServersConfig struct {
	CustomHeaders map[string]interface{} `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// example:
	//
	// SSE
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// example:
	//
	// https://example.com/mcp/sse
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListMcpServersResponseBodyPagingInfoMcpServersConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyPagingInfoMcpServersConfig) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) GetCustomHeaders() map[string]interface{} {
	return s.CustomHeaders
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) GetTransport() *string {
	return s.Transport
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) GetUrl() *string {
	return s.Url
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) SetCustomHeaders(v map[string]interface{}) *ListMcpServersResponseBodyPagingInfoMcpServersConfig {
	s.CustomHeaders = v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) SetTransport(v string) *ListMcpServersResponseBodyPagingInfoMcpServersConfig {
	s.Transport = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) SetUrl(v string) *ListMcpServersResponseBodyPagingInfoMcpServersConfig {
	s.Url = &v
	return s
}

func (s *ListMcpServersResponseBodyPagingInfoMcpServersConfig) Validate() error {
	return dara.Validate(s)
}
