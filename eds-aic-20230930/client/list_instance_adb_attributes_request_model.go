// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAdbAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *ListInstanceAdbAttributesRequest
	GetExternalIp() *string
	SetInstanceIds(v []*string) *ListInstanceAdbAttributesRequest
	GetInstanceIds() []*string
	SetInternalIp(v string) *ListInstanceAdbAttributesRequest
	GetInternalIp() *string
	SetInternalPort(v string) *ListInstanceAdbAttributesRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *ListInstanceAdbAttributesRequest
	GetIpProtocol() *string
	SetMaxResults(v int32) *ListInstanceAdbAttributesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceAdbAttributesRequest
	GetNextToken() *string
}

type ListInstanceAdbAttributesRequest struct {
	// example:
	//
	// 106.38.188.223
	ExternalIp  *string   `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 10.0.3.23
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// example:
	//
	// 5555
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstanceAdbAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAdbAttributesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAdbAttributesRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *ListInstanceAdbAttributesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListInstanceAdbAttributesRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *ListInstanceAdbAttributesRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *ListInstanceAdbAttributesRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListInstanceAdbAttributesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceAdbAttributesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceAdbAttributesRequest) SetExternalIp(v string) *ListInstanceAdbAttributesRequest {
	s.ExternalIp = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetInstanceIds(v []*string) *ListInstanceAdbAttributesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetInternalIp(v string) *ListInstanceAdbAttributesRequest {
	s.InternalIp = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetInternalPort(v string) *ListInstanceAdbAttributesRequest {
	s.InternalPort = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetIpProtocol(v string) *ListInstanceAdbAttributesRequest {
	s.IpProtocol = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetMaxResults(v int32) *ListInstanceAdbAttributesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) SetNextToken(v string) *ListInstanceAdbAttributesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceAdbAttributesRequest) Validate() error {
	return dara.Validate(s)
}
