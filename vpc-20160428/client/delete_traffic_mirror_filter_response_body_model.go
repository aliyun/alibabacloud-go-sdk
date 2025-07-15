// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficMirrorFilterResponseBody
	GetRequestId() *string
}

type DeleteTrafficMirrorFilterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 07F272E2-6AD5-433A-8207-A607C76F1676
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficMirrorFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficMirrorFilterResponseBody) SetRequestId(v string) *DeleteTrafficMirrorFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
