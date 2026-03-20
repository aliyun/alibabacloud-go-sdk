// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupServersRequest
	GetNextToken() *string
	SetServerGroupId(v string) *ListServerGroupServersRequest
	GetServerGroupId() *string
	SetServerIds(v []*string) *ListServerGroupServersRequest
	GetServerIds() []*string
	SetServerIps(v []*string) *ListServerGroupServersRequest
	GetServerIps() []*string
	SetSkip(v int32) *ListServerGroupServersRequest
	GetSkip() *int32
}

type ListServerGroupServersRequest struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server IDs.
	//
	// You can specify at most 200 servers in each call.
	ServerIds []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	// The server IP addresses.
	//
	// You can specify at most 200 servers in each call.
	ServerIps []*string `json:"ServerIps,omitempty" xml:"ServerIps,omitempty" type:"Repeated"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListServerGroupServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupServersRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListServerGroupServersRequest) GetServerIds() []*string {
	return s.ServerIds
}

func (s *ListServerGroupServersRequest) GetServerIps() []*string {
	return s.ServerIps
}

func (s *ListServerGroupServersRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIps(v []*string) *ListServerGroupServersRequest {
	s.ServerIps = v
	return s
}

func (s *ListServerGroupServersRequest) SetSkip(v int32) *ListServerGroupServersRequest {
	s.Skip = &v
	return s
}

func (s *ListServerGroupServersRequest) Validate() error {
	return dara.Validate(s)
}
