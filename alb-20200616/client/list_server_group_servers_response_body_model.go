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
	// The maximum number of entries returned.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If **NextToken*	- is not empty, the value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of backend servers.
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
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
	// The description of the backend server.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether the remote IP address feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RemoteIpEnabled *bool `json:"RemoteIpEnabled,omitempty" xml:"RemoteIpEnabled,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// sgp-qy042e1jabmprh****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The ID of the backend server.
	//
	// > If **ServerType*	- is set to **Fc**, **ServerId*	- is the ARN of a function.
	//
	// example:
	//
	// i-bp1f9kdprbgy9uiu****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address in inclusive ENI mode.
	//
	// example:
	//
	// 192.168.XX.XX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The status of the backend server. Valid values:
	//
	// 	- **Adding**
	//
	// 	- **Available**
	//
	// 	- **Configuring**
	//
	// 	- **Removing**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The weight of the backend server. An ECS instance with a higher weight receives more requests.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) GetDescription() *string {
	return s.Description
}

func (s *ListServerGroupServersResponseBodyServers) GetPort() *int32 {
	return s.Port
}

func (s *ListServerGroupServersResponseBodyServers) GetRemoteIpEnabled() *bool {
	return s.RemoteIpEnabled
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

func (s *ListServerGroupServersResponseBodyServers) GetWeight() *int32 {
	return s.Weight
}

func (s *ListServerGroupServersResponseBodyServers) SetDescription(v string) *ListServerGroupServersResponseBodyServers {
	s.Description = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetRemoteIpEnabled(v bool) *ListServerGroupServersResponseBodyServers {
	s.RemoteIpEnabled = &v
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

func (s *ListServerGroupServersResponseBodyServers) SetWeight(v int32) *ListServerGroupServersResponseBodyServers {
	s.Weight = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) Validate() error {
	return dara.Validate(s)
}
