// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutoShutdownPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryAutoShutdownPoliciesRequest
	GetInstanceId() *string
	SetToken(v string) *QueryAutoShutdownPoliciesRequest
	GetToken() *string
	SetInstanceIds(v []*string) *QueryAutoShutdownPoliciesRequest
	GetInstanceIds() []*string
}

type QueryAutoShutdownPoliciesRequest struct {
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// *******
	Token       *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s QueryAutoShutdownPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAutoShutdownPoliciesRequest) GoString() string {
	return s.String()
}

func (s *QueryAutoShutdownPoliciesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryAutoShutdownPoliciesRequest) GetToken() *string {
	return s.Token
}

func (s *QueryAutoShutdownPoliciesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *QueryAutoShutdownPoliciesRequest) SetInstanceId(v string) *QueryAutoShutdownPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryAutoShutdownPoliciesRequest) SetToken(v string) *QueryAutoShutdownPoliciesRequest {
	s.Token = &v
	return s
}

func (s *QueryAutoShutdownPoliciesRequest) SetInstanceIds(v []*string) *QueryAutoShutdownPoliciesRequest {
	s.InstanceIds = v
	return s
}

func (s *QueryAutoShutdownPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
