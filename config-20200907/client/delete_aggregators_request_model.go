// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorIds(v string) *DeleteAggregatorsRequest
	GetAggregatorIds() *string
	SetClientToken(v string) *DeleteAggregatorsRequest
	GetClientToken() *string
}

type DeleteAggregatorsRequest struct {
	// The ID of the account group. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-9190626622af00a9****
	AggregatorIds *string `json:"AggregatorIds,omitempty" xml:"AggregatorIds,omitempty"`
	// The client token that you want to use to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteAggregatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregatorsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsRequest) GetAggregatorIds() *string {
	return s.AggregatorIds
}

func (s *DeleteAggregatorsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAggregatorsRequest) SetAggregatorIds(v string) *DeleteAggregatorsRequest {
	s.AggregatorIds = &v
	return s
}

func (s *DeleteAggregatorsRequest) SetClientToken(v string) *DeleteAggregatorsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAggregatorsRequest) Validate() error {
	return dara.Validate(s)
}
