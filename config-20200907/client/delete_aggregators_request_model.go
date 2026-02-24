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
	// This parameter is required.
	AggregatorIds *string `json:"AggregatorIds,omitempty" xml:"AggregatorIds,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
