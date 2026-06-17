// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectPolarClawDevicePairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RejectPolarClawDevicePairRequest
	GetApplicationId() *string
	SetPairRequestId(v string) *RejectPolarClawDevicePairRequest
	GetPairRequestId() *string
}

type RejectPolarClawDevicePairRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The pairing request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-abc-123
	PairRequestId *string `json:"PairRequestId,omitempty" xml:"PairRequestId,omitempty"`
}

func (s RejectPolarClawDevicePairRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectPolarClawDevicePairRequest) GoString() string {
	return s.String()
}

func (s *RejectPolarClawDevicePairRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RejectPolarClawDevicePairRequest) GetPairRequestId() *string {
	return s.PairRequestId
}

func (s *RejectPolarClawDevicePairRequest) SetApplicationId(v string) *RejectPolarClawDevicePairRequest {
	s.ApplicationId = &v
	return s
}

func (s *RejectPolarClawDevicePairRequest) SetPairRequestId(v string) *RejectPolarClawDevicePairRequest {
	s.PairRequestId = &v
	return s
}

func (s *RejectPolarClawDevicePairRequest) Validate() error {
	return dara.Validate(s)
}
