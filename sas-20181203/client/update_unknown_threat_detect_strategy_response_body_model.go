// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUnknownThreatDetectStrategyResponseBody
	GetRequestId() *string
}

type UpdateUnknownThreatDetectStrategyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUnknownThreatDetectStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUnknownThreatDetectStrategyResponseBody) SetRequestId(v string) *UpdateUnknownThreatDetectStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
