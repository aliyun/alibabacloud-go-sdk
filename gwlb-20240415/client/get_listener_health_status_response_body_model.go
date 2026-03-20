// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody
	GetListenerHealthStatus() []*GetListenerHealthStatusResponseBodyListenerHealthStatus
	SetMaxResults(v int32) *GetListenerHealthStatusResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetListenerHealthStatusResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetListenerHealthStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetListenerHealthStatusResponseBody
	GetTotalCount() *int32
}

type GetListenerHealthStatusResponseBody struct {
	// The health check status of the server groups that are associated with the listener.
	ListenerHealthStatus []*GetListenerHealthStatusResponseBodyListenerHealthStatus `json:"ListenerHealthStatus,omitempty" xml:"ListenerHealthStatus,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// U12WEI6Ro2ol3wA54rBNSwdC5+lYy6q5SjIQEvc1wz5mjZxV+YjsHRdXV8XauY1BpOQIvwX63E0en54H3D****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED4F222-36A0-5470-8A9A-AAB4E96BAC1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetListenerHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBody) GetListenerHealthStatus() []*GetListenerHealthStatusResponseBodyListenerHealthStatus {
	return s.ListenerHealthStatus
}

func (s *GetListenerHealthStatusResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetListenerHealthStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetListenerHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListenerHealthStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetListenerHealthStatusResponseBody) SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody {
	s.ListenerHealthStatus = v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetMaxResults(v int32) *GetListenerHealthStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetNextToken(v string) *GetListenerHealthStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetRequestId(v string) *GetListenerHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetTotalCount(v int32) *GetListenerHealthStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) Validate() error {
	if s.ListenerHealthStatus != nil {
		for _, item := range s.ListenerHealthStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyListenerHealthStatus struct {
	// The listener ID.
	//
	// example:
	//
	// lsn-sg8aha6pzjavvo****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The information about the server groups.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) GetServerGroupInfos() []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	return s.ServerGroupInfos
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ServerGroupInfos = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) Validate() error {
	if s.ServerGroupInfos != nil {
		for _, item := range s.ServerGroupInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos struct {
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-0vdsbyszro3nr6****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers.
	Servers []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetServers() []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	return s.Servers
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetHealthCheckEnabled(v bool) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServers(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.Servers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) Validate() error {
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

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers struct {
	// The backend port.
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The reason why **Status*	- indicates an unhealthy status.
	Reason *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The backend server ID.
	//
	// example:
	//
	// i-2ze4rnh8yj9kif3z****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 192.168.0.XXX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The health status of the backend server. Valid values:
	//
	// 	- **Initial**: Health checks are configured for the GWLB instance, but no data is found.
	//
	// 	- **Unhealthy**: The backend server consecutively fails health checks.
	//
	// 	- **Unused**: The backend server is not in use.
	//
	// 	- **Unavailable**: Health checks are disabled.
	//
	// 	- **Healthy**: The backend server is healthy.
	//
	// example:
	//
	// Healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GetPort() *int32 {
	return s.Port
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GetReason() *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason {
	return s.Reason
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GetServerId() *string {
	return s.ServerId
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GetStatus() *string {
	return s.Status
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetReason(v *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Status = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason struct {
	// The reason why **Status*	- indicates an unhealthy status. Valid values:
	//
	// 	- **CONNECT_TIMEOUT**: The GWLB instance failed to connect to the backend server within the specified period of time.
	//
	// 	- **CONNECT_FAILED**: The GWLB instance failed to connect to the backend server.
	//
	// 	- **RECV_RESPONSE_TIMEOUT**: The GWLB instance failed to receive a response from the backend server within the specified period of time.
	//
	// 	- **CONNECT_INTERRUPT**: The connection between the health check and the backend server was interrupted.
	//
	// 	- **HTTP_CODE_NOT_MATCH**: The HTTP status code from the backend server is not the expected one.
	//
	// 	- **HTTP_INVALID_HEADER**: The format of the response from the backend servers is invalid.
	//
	// example:
	//
	// CONNECT_TIMEOUT
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason {
	s.ReasonCode = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) Validate() error {
	return dara.Validate(s)
}
