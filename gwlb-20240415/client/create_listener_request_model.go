// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateListenerRequest
	GetDryRun() *bool
	SetListenerDescription(v string) *CreateListenerRequest
	GetListenerDescription() *string
	SetLoadBalancerId(v string) *CreateListenerRequest
	GetLoadBalancerId() *string
	SetServerGroupId(v string) *CreateListenerRequest
	GetServerGroupId() *string
	SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest
	GetTag() []*CreateListenerRequestTag
	SetTcpIdleTimeout(v int32) *CreateListenerRequest
	GetTcpIdleTimeout() *int32
}

type CreateListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the listener.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener-description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-te609d6696632f7*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-ckh01px70dszof****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag []*CreateListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The timeout period of an idle TCP connection. Unit: seconds.
	//
	// Default value: **350**.
	//
	// Valid values: **60*	- to **6000**.
	//
	// example:
	//
	// 350
	TcpIdleTimeout *int32 `json:"TcpIdleTimeout,omitempty" xml:"TcpIdleTimeout,omitempty"`
}

func (s CreateListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateListenerRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *CreateListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateListenerRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateListenerRequest) GetTag() []*CreateListenerRequestTag {
	return s.Tag
}

func (s *CreateListenerRequest) GetTcpIdleTimeout() *int32 {
	return s.TcpIdleTimeout
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetServerGroupId(v string) *CreateListenerRequest {
	s.ServerGroupId = &v
	return s
}

func (s *CreateListenerRequest) SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateListenerRequest) SetTcpIdleTimeout(v int32) *CreateListenerRequest {
	s.TcpIdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerRequestTag struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateListenerRequestTag) SetKey(v string) *CreateListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateListenerRequestTag) SetValue(v string) *CreateListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateListenerRequestTag) Validate() error {
	return dara.Validate(s)
}
