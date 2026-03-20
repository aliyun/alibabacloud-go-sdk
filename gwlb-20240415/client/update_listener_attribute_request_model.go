// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateListenerAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateListenerAttributeRequest
	GetDryRun() *bool
	SetListenerDescription(v string) *UpdateListenerAttributeRequest
	GetListenerDescription() *string
	SetListenerId(v string) *UpdateListenerAttributeRequest
	GetListenerId() *string
	SetServerGroupId(v string) *UpdateListenerAttributeRequest
	GetServerGroupId() *string
	SetTcpIdleTimeout(v int32) *UpdateListenerAttributeRequest
	GetTcpIdleTimeout() *int32
}

type UpdateListenerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener description.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener_description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-lxce8iqbof2vl0****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-sp8d2r6y7t0xtl****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The timeout period of an idle TCP connection. Unit: seconds.
	//
	// Valid values: **60*	- to **6000**.
	//
	// example:
	//
	// 350
	TcpIdleTimeout *int32 `json:"TcpIdleTimeout,omitempty" xml:"TcpIdleTimeout,omitempty"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateListenerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateListenerAttributeRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *UpdateListenerAttributeRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateListenerAttributeRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateListenerAttributeRequest) GetTcpIdleTimeout() *int32 {
	return s.TcpIdleTimeout
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetServerGroupId(v string) *UpdateListenerAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetTcpIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.TcpIdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
