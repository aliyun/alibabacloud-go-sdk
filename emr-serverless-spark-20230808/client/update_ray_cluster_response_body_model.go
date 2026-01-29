// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateRayClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpdateRayClusterResponseBody
	GetRequestId() *string
}

type UpdateRayClusterResponseBody struct {
	// example:
	//
	// ray-xxxxxxxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRayClusterResponseBody) SetClusterId(v string) *UpdateRayClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateRayClusterResponseBody) SetRequestId(v string) *UpdateRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRayClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
