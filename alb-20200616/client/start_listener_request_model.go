// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *StartListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *StartListenerRequest
	GetListenerId() *string
}

type StartListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the **DryRunOperation*	- error code is returned.
	//
	// 	- **false**: (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Application Load Balancer (ALB) listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s StartListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartListenerRequest) GoString() string {
	return s.String()
}

func (s *StartListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *StartListenerRequest) SetClientToken(v string) *StartListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StartListenerRequest) SetDryRun(v bool) *StartListenerRequest {
	s.DryRun = &v
	return s
}

func (s *StartListenerRequest) SetListenerId(v string) *StartListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *StartListenerRequest) Validate() error {
	return dara.Validate(s)
}
