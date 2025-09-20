// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationDateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLocationDateClusterResponseBody
	GetRequestId() *string
}

type UpdateLocationDateClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52B017A8-FEF5-0A61-BAEE-234A8AD8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLocationDateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationDateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLocationDateClusterResponseBody) SetRequestId(v string) *UpdateLocationDateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLocationDateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
