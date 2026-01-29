// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateRayClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateRayClusterResponseBody
	GetRequestId() *string
}

type CreateRayClusterResponseBody struct {
	// example:
	//
	// ray-k7nm8ahl5te4tg91
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRayClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRayClusterResponseBody) SetClusterId(v string) *CreateRayClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateRayClusterResponseBody) SetRequestId(v string) *CreateRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRayClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
