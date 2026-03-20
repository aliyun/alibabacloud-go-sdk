// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupServersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupServersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServerGroupServersResponseBody
	GetRequestId() *string
	SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody
	GetServers() []*ListServerGroupServersResponseBodyServers
	SetTotalCount(v int32) *ListServerGroupServersResponseBody
	GetTotalCount() *int32
}

type ListServerGroupServersResponseBody struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The backend servers.
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupServersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerGroupServersResponseBody) GetServers() []*ListServerGroupServersResponseBodyServers {
	return s.Servers
}

func (s *ListServerGroupServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServerGroupServersResponseBody) SetMaxResults(v int32) *ListServerGroupServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetNextToken(v string) *ListServerGroupServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetTotalCount(v int32) *ListServerGroupServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServerGroupServersResponseBody) Validate() error {
	if s.Servers != nil {
		for _, item := range s.Servers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerGroupServersResponseBodyServers struct {
	// The backend server port. Valid values:
	//
	// 	- **6081**
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend server ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.xxx.xxx
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: elastic network interface (ENI)
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// Indicates the status of the backend server. Valid values:
	//
	// 	- **Adding**: The backend server is being added.
	//
	// 	- **Available**: The backend server is available.
	//
	// 	- **Draining**: The backend server is in connection draining.
	//
	// 	- **Removing**: The backend server is being removed.
	//
	// 	- **Replacing**: The backend server is being replaced.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) GetPort() *int32 {
	return s.Port
}

func (s *ListServerGroupServersResponseBodyServers) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListServerGroupServersResponseBodyServers) GetServerId() *string {
	return s.ServerId
}

func (s *ListServerGroupServersResponseBodyServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *ListServerGroupServersResponseBodyServers) GetServerType() *string {
	return s.ServerType
}

func (s *ListServerGroupServersResponseBodyServers) GetStatus() *string {
	return s.Status
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerGroupId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerIp(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerType(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetStatus(v string) *ListServerGroupServersResponseBodyServers {
	s.Status = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) Validate() error {
	return dara.Validate(s)
}
