// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficMirrorSessionResponseBody
	GetRequestId() *string
}

type DeleteTrafficMirrorSessionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficMirrorSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficMirrorSessionResponseBody) SetRequestId(v string) *DeleteTrafficMirrorSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficMirrorSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
