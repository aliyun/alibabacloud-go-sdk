// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnknownThreatDetectStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateUnknownThreatDetectStrategyResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateUnknownThreatDetectStrategyResponseBody
	GetRequestId() *string
}

type CreateUnknownThreatDetectStrategyResponseBody struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUnknownThreatDetectStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUnknownThreatDetectStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUnknownThreatDetectStrategyResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateUnknownThreatDetectStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUnknownThreatDetectStrategyResponseBody) SetId(v int64) *CreateUnknownThreatDetectStrategyResponseBody {
	s.Id = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyResponseBody) SetRequestId(v string) *CreateUnknownThreatDetectStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
