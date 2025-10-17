// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterEndpointZonalResponseBody
	GetRequestId() *string
}

type ModifyDBClusterEndpointZonalResponseBody struct {
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterEndpointZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterEndpointZonalResponseBody) SetRequestId(v string) *ModifyDBClusterEndpointZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
