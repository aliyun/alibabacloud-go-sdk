// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficMirrorFilterRulesResponseBody
	GetRequestId() *string
}

type DeleteTrafficMirrorFilterRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BD8A3F71-00C5-4655-8F55-11F3976C3274
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficMirrorFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficMirrorFilterRulesResponseBody) SetRequestId(v string) *DeleteTrafficMirrorFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
