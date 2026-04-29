// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUnknownThreatDetectProcessResponseBody
	GetRequestId() *string
}

type UpdateUnknownThreatDetectProcessResponseBody struct {
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUnknownThreatDetectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUnknownThreatDetectProcessResponseBody) SetRequestId(v string) *UpdateUnknownThreatDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUnknownThreatDetectProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
