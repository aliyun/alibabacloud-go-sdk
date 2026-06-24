// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaPvlNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointName(v string) *UpdateKibanaPvlNetworkRequest
	GetEndpointName() *string
	SetSecurityGroups(v []*string) *UpdateKibanaPvlNetworkRequest
	GetSecurityGroups() []*string
	SetClientToken(v string) *UpdateKibanaPvlNetworkRequest
	GetClientToken() *string
	SetPvlId(v string) *UpdateKibanaPvlNetworkRequest
	GetPvlId() *string
}

type UpdateKibanaPvlNetworkRequest struct {
	// The endpoint name.
	//
	// example:
	//
	// es-cn-text-kibana
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	// The list of security groups.
	SecurityGroups []*string `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" type:"Repeated"`
	// Used to ensure the idempotency of the request.
	//
	// example:
	//
	// xxxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The Kibana private network connection ID.
	//
	// example:
	//
	// es-cn-vo93ngti8000a****-kibana-internal-internal
	PvlId *string `json:"pvlId,omitempty" xml:"pvlId,omitempty"`
}

func (s UpdateKibanaPvlNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaPvlNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKibanaPvlNetworkRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateKibanaPvlNetworkRequest) GetSecurityGroups() []*string {
	return s.SecurityGroups
}

func (s *UpdateKibanaPvlNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateKibanaPvlNetworkRequest) GetPvlId() *string {
	return s.PvlId
}

func (s *UpdateKibanaPvlNetworkRequest) SetEndpointName(v string) *UpdateKibanaPvlNetworkRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateKibanaPvlNetworkRequest) SetSecurityGroups(v []*string) *UpdateKibanaPvlNetworkRequest {
	s.SecurityGroups = v
	return s
}

func (s *UpdateKibanaPvlNetworkRequest) SetClientToken(v string) *UpdateKibanaPvlNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKibanaPvlNetworkRequest) SetPvlId(v string) *UpdateKibanaPvlNetworkRequest {
	s.PvlId = &v
	return s
}

func (s *UpdateKibanaPvlNetworkRequest) Validate() error {
	return dara.Validate(s)
}
