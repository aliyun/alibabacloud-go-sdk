// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromMountPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachVscFromMountPointsResponseBody
	GetRequestId() *string
}

type DetachVscFromMountPointsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVscFromMountPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVscFromMountPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachVscFromMountPointsResponseBody) SetRequestId(v string) *DetachVscFromMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachVscFromMountPointsResponseBody) Validate() error {
	return dara.Validate(s)
}
