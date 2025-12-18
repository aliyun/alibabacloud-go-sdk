// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateAggregatorsResult(v *DeleteAggregatorsResponseBodyOperateAggregatorsResult) *DeleteAggregatorsResponseBody
	GetOperateAggregatorsResult() *DeleteAggregatorsResponseBodyOperateAggregatorsResult
	SetRequestId(v string) *DeleteAggregatorsResponseBody
	GetRequestId() *string
}

type DeleteAggregatorsResponseBody struct {
	// The returned result.
	OperateAggregatorsResult *DeleteAggregatorsResponseBodyOperateAggregatorsResult `json:"OperateAggregatorsResult,omitempty" xml:"OperateAggregatorsResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8195B664-9565-4685-89AC-8B5F04B44B92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBody) GetOperateAggregatorsResult() *DeleteAggregatorsResponseBodyOperateAggregatorsResult {
	return s.OperateAggregatorsResult
}

func (s *DeleteAggregatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggregatorsResponseBody) SetOperateAggregatorsResult(v *DeleteAggregatorsResponseBodyOperateAggregatorsResult) *DeleteAggregatorsResponseBody {
	s.OperateAggregatorsResult = v
	return s
}

func (s *DeleteAggregatorsResponseBody) SetRequestId(v string) *DeleteAggregatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggregatorsResponseBody) Validate() error {
	if s.OperateAggregatorsResult != nil {
		if err := s.OperateAggregatorsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAggregatorsResponseBodyOperateAggregatorsResult struct {
	// The details of the account group.
	OperateAggregators []*DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators `json:"OperateAggregators,omitempty" xml:"OperateAggregators,omitempty" type:"Repeated"`
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResult) GetOperateAggregators() []*DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	return s.OperateAggregators
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResult) SetOperateAggregators(v []*DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) *DeleteAggregatorsResponseBodyOperateAggregatorsResult {
	s.OperateAggregators = v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResult) Validate() error {
	if s.OperateAggregators != nil {
		for _, item := range s.OperateAggregators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators struct {
	// The ID of the account group.
	//
	// example:
	//
	// ca-dacf86d8314e00eb****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The error code returned.
	//
	// > No error code is returned for the account group if the account group is deleted.
	//
	// example:
	//
	// AccountNotExisted
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetAggregatorId(v string) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetErrorCode(v string) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetSuccess(v bool) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.Success = &v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) Validate() error {
	return dara.Validate(s)
}
