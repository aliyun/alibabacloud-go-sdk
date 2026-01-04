// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveRouteTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchActiveRouteTargetResponseBody
	GetRequestId() *string
}

type SwitchActiveRouteTargetResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// C1221A1F-2ACD-4592-8F27-474E02883159
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchActiveRouteTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveRouteTargetResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchActiveRouteTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchActiveRouteTargetResponseBody) SetRequestId(v string) *SwitchActiveRouteTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchActiveRouteTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
