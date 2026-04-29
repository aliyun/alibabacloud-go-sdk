// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUnknownThreatDetectStrategyResponseBody
	GetRequestId() *string
}

type DeleteUnknownThreatDetectStrategyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUnknownThreatDetectStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUnknownThreatDetectStrategyResponseBody) SetRequestId(v string) *DeleteUnknownThreatDetectStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUnknownThreatDetectStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
