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
	SetNextToken(v string) *GetListenerHealthStatusResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetListenerHealthStatusResponseBody
	GetRequestId() *string
	SetRuleHealthStatus(v []*GetListenerHealthStatusResponseBodyRuleHealthStatus) *GetListenerHealthStatusResponseBody
	GetRuleHealthStatus() []*GetListenerHealthStatusResponseBodyRuleHealthStatus
}

type GetListenerHealthStatusResponseBody struct {
	// The health check status of the server groups associated with the listener.
	ListenerHealthStatus []*GetListenerHealthStatusResponseBodyListenerHealthStatus `json:"ListenerHealthStatus,omitempty" xml:"ListenerHealthStatus,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The health check status of the forwarding rules.
	RuleHealthStatus []*GetListenerHealthStatusResponseBodyRuleHealthStatus `json:"RuleHealthStatus,omitempty" xml:"RuleHealthStatus,omitempty" type:"Repeated"`
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

func (s *GetListenerHealthStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetListenerHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListenerHealthStatusResponseBody) GetRuleHealthStatus() []*GetListenerHealthStatusResponseBodyRuleHealthStatus {
	return s.RuleHealthStatus
}

func (s *GetListenerHealthStatusResponseBody) SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody {
	s.ListenerHealthStatus = v
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

func (s *GetListenerHealthStatusResponseBody) SetRuleHealthStatus(v []*GetListenerHealthStatusResponseBodyRuleHealthStatus) *GetListenerHealthStatusResponseBody {
	s.RuleHealthStatus = v
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
	if s.RuleHealthStatus != nil {
		for _, item := range s.RuleHealthStatus {
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
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol.
	//
	// example:
	//
	// http
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The information about the server group.
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

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) GetServerGroupInfos() []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	return s.ServerGroupInfos
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerProtocol(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerProtocol = &v
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
	// The action specified for the server group. Valid values:
	//
	// 	- **ForwardGroup**: distributes requests to server groups.
	//
	// 	- **TrafficMirror**: mirrors requests to server groups.
	//
	// example:
	//
	// TrafficMirror
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Indicates whether health checks are enabled. If **on*	- is returned, it indicates that health checks are enabled.
	//
	// example:
	//
	// on
	HealthCheckEnabled *string `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// A list of unhealthy backend servers.
	NonNormalServers []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers `json:"NonNormalServers,omitempty" xml:"NonNormalServers,omitempty" type:"Repeated"`
	// The ID of the server group that is associated with the listener.
	//
	// example:
	//
	// vsp-bp1qjwo61pqz3ahltv****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetActionType() *string {
	return s.ActionType
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetHealthCheckEnabled() *string {
	return s.HealthCheckEnabled
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetNonNormalServers() []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	return s.NonNormalServers
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetActionType(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ActionType = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetHealthCheckEnabled(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetNonNormalServers(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.NonNormalServers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) Validate() error {
	if s.NonNormalServers != nil {
		for _, item := range s.NonNormalServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers struct {
	// The backend port.
	//
	// example:
	//
	// 90
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The cause for the unhealthy state of the backend servers.
	Reason *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The ID of the backend server.
	//
	// example:
	//
	// rg-bp1bfa08ex*****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.8.10
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The status of the health check. Valid values: Valid values:
	//
	// 	- **Initial**: indicates that health checks are configured for the NLB instance, but no data was found.
	//
	// 	- **Unhealthy**: indicates that the backend server consecutively fails health checks.
	//
	// 	- **Unused**: indicates that the weight of the backend server is 0.
	//
	// 	- **Unavailable**: indicates that health checks are disabled.
	//
	// example:
	//
	// Initial
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GetPort() *int32 {
	return s.Port
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GetReason() *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	return s.Reason
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GetServerId() *string {
	return s.ServerId
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GetStatus() *string {
	return s.Status
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetReason(v *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Status = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason struct {
	// The HTTP status code returned from the server, for example, **302**.
	//
	// > A value is returned only if `ReasonCode` is set to **RESPONSE_MISMATCH**.
	//
	// example:
	//
	// 302
	ActualResponse *string `json:"ActualResponse,omitempty" xml:"ActualResponse,omitempty"`
	// The HTTP status code returned after backend servers pass health checks.
	//
	// Valid values: **HTTP_2xx**, **HTTP_3xx**, **HTTP_4xx**, and **HTTP_5xx**. Multiple status codes are separated by commas (,).
	//
	// > This value is returned only if **ReasonCode*	- is set to **RESPONSE_MISMATCH**.
	//
	// example:
	//
	// HTTP_2xx
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// The reason why the value of **Status*	- is Unhealthy. Only HTTP and HTTPS listeners support this parameter.
	//
	// 	- **CONNECT_TIMEOUT**: ALB failed to connect to the backend server within the specified period of time.
	//
	// 	- **CONNECT_FAILED**: ALB failed to connect to the backend server.
	//
	// 	- **RECV_RESPONSE_FAILED**: ALB failed to receive a response from the backend server.
	//
	// 	- **RECV_RESPONSE_TIMEOUT**: ALB failed to receive a response from the backend server within the specified period of time.
	//
	// 	- **SEND_REQUEST_FAILED**: ALB failed to send a request to the backend server.
	//
	// 	- **SEND_REQUEST_TIMEOUT**: ALB failed to send a request to the backend server within the specified period of time.
	//
	// 	- **RESPONSE_FORMAT_ERROR**: The format of the response from the backend server is invalid.
	//
	// 	- **RESPONSE_MISMATCH**: The HTTP status code returned from the backend server is not the expected one.
	//
	// example:
	//
	// RESPONSE_MISMATCH
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GetActualResponse() *string {
	return s.ActualResponse
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GetExpectedResponse() *string {
	return s.ExpectedResponse
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetActualResponse(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ActualResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetExpectedResponse(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ExpectedResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ReasonCode = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) Validate() error {
	return dara.Validate(s)
}

type GetListenerHealthStatusResponseBodyRuleHealthStatus struct {
	// The ID of the forwarding rule.
	//
	// example:
	//
	// rule-hp34s2h0xx1ht4nwo****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The server groups.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatus) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) GetRuleId() *string {
	return s.RuleId
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) GetServerGroupInfos() []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	return s.ServerGroupInfos
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) SetRuleId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatus {
	s.RuleId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyRuleHealthStatus {
	s.ServerGroupInfos = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatus) Validate() error {
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

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos struct {
	// The action specified for the server group.
	//
	// example:
	//
	// TrafficMirror
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Indicates whether health checks are enabled. If **on*	- is returned, it indicates that health checks are enabled.
	//
	// example:
	//
	// on
	HealthCheckEnabled *string `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// A list of unhealthy backend servers.
	NonNormalServers []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers `json:"NonNormalServers,omitempty" xml:"NonNormalServers,omitempty" type:"Repeated"`
	// The ID of the server group that is associated with the listener.
	//
	// example:
	//
	// vsp-bp1qjwo61pqz3ahlt****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GetActionType() *string {
	return s.ActionType
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GetHealthCheckEnabled() *string {
	return s.HealthCheckEnabled
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GetNonNormalServers() []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	return s.NonNormalServers
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetActionType(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.ActionType = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetHealthCheckEnabled(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetNonNormalServers(v []*GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.NonNormalServers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfos) Validate() error {
	if s.NonNormalServers != nil {
		for _, item := range s.NonNormalServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers struct {
	// The backend port.
	//
	// example:
	//
	// 90
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The cause for the unhealthy state of the backend servers.
	Reason *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The ID of the backend server.
	//
	// example:
	//
	// rg-bp1bfa08ex****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server group.
	//
	// example:
	//
	// 192.168.2.11
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The status of the health check. Valid values: Valid values:
	//
	// 	- **Initial**: indicates that health checks are configured for the NLB instance, but no data was found.
	//
	// 	- **Unhealthy**: indicates that the backend server consecutively fails health checks.
	//
	// 	- **Unused**: indicates that the weight of the backend server is 0.
	//
	// 	- **Unavailable**: indicates that health checks are disabled.
	//
	// example:
	//
	// Initial
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GetPort() *int32 {
	return s.Port
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GetReason() *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	return s.Reason
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GetServerId() *string {
	return s.ServerId
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) GetStatus() *string {
	return s.Status
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetReason(v *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers {
	s.Status = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServers) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason struct {
	// The HTTP status code returned from the server, for example, **302**.
	//
	// > A value is returned only if **ReasonCode*	- is set to **RESPONSE_MISMATCH**.
	//
	// example:
	//
	// 302
	ActualResponse *string `json:"ActualResponse,omitempty" xml:"ActualResponse,omitempty"`
	// The HTTP status code returned after backend servers pass health checks.
	//
	// Valid values: **HTTP_2xx**, **HTTP_3xx**, **HTTP_4xx**, and **HTTP_5xx**. Multiple status codes are separated by commas (,).
	//
	// > A value is returned only if **ReasonCode*	- is set to **RESPONSE_MISMATCH**.
	//
	// example:
	//
	// HTTP_2xx
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// The reason why the value of **Status*	- is Unhealthy. Only forwarding rules for HTTP and HTTPS listeners support this parameter.
	//
	// 	- **CONNECT_TIMEOUT**: ALB failed to connect to the backend server within the specified period of time.
	//
	// 	- **CONNECT_FAILED**: ALB failed to connect to the backend server.
	//
	// 	- **RECV_RESPONSE_FAILED**: ALB failed to receive a response from the backend server.
	//
	// 	- **RECV_RESPONSE_TIMEOUT**: ALB failed to receive a response from the backend server within the specified period of time.
	//
	// 	- **SEND_REQUEST_FAILED**: ALB failed to send a request to the backend server.
	//
	// 	- **SEND_REQUEST_TIMEOUT**: ALB failed to send a request to the backend server within the specified period of time.
	//
	// 	- **RESPONSE_FORMAT_ERROR**: The format of the response from the backend server is invalid.
	//
	// 	- **RESPONSE_MISMATCH**: The HTTP status code returned from the backend server is not the expected one.
	//
	// example:
	//
	// RESPONSE_MISMATCH
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) GetActualResponse() *string {
	return s.ActualResponse
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) GetExpectedResponse() *string {
	return s.ExpectedResponse
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetActualResponse(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ActualResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetExpectedResponse(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ExpectedResponse = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason {
	s.ReasonCode = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyRuleHealthStatusServerGroupInfosNonNormalServersReason) Validate() error {
	return dara.Validate(s)
}
