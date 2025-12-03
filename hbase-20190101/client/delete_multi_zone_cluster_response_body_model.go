// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiZoneClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMultiZoneClusterResponseBody
	GetRequestId() *string
}

type DeleteMultiZoneClusterResponseBody struct {
	// example:
	//
	// 169A3910-A39E-4BC2-AA9F-E7AD8D473527
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultiZoneClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultiZoneClusterResponseBody) SetRequestId(v string) *DeleteMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultiZoneClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
