// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceMeshResponseBody
	GetRequestId() *string
}

type DeleteServiceMeshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceMeshResponseBody) SetRequestId(v string) *DeleteServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
