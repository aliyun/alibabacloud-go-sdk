// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePolarClawDevicePairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ApprovePolarClawDevicePairRequest
	GetApplicationId() *string
	SetPairRequestId(v string) *ApprovePolarClawDevicePairRequest
	GetPairRequestId() *string
}

type ApprovePolarClawDevicePairRequest struct {
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

func (s ApprovePolarClawDevicePairRequest) String() string {
	return dara.Prettify(s)
}

func (s ApprovePolarClawDevicePairRequest) GoString() string {
	return s.String()
}

func (s *ApprovePolarClawDevicePairRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ApprovePolarClawDevicePairRequest) GetPairRequestId() *string {
	return s.PairRequestId
}

func (s *ApprovePolarClawDevicePairRequest) SetApplicationId(v string) *ApprovePolarClawDevicePairRequest {
	s.ApplicationId = &v
	return s
}

func (s *ApprovePolarClawDevicePairRequest) SetPairRequestId(v string) *ApprovePolarClawDevicePairRequest {
	s.PairRequestId = &v
	return s
}

func (s *ApprovePolarClawDevicePairRequest) Validate() error {
	return dara.Validate(s)
}
