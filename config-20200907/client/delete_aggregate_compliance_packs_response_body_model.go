// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateCompliancePacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateCompliancePacksResult(v *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteAggregateCompliancePacksResponseBody
	GetOperateCompliancePacksResult() *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult
	SetRequestId(v string) *DeleteAggregateCompliancePacksResponseBody
	GetRequestId() *string
}

type DeleteAggregateCompliancePacksResponseBody struct {
	OperateCompliancePacksResult *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult `json:"OperateCompliancePacksResult,omitempty" xml:"OperateCompliancePacksResult,omitempty" type:"Struct"`
	RequestId                    *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateCompliancePacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBody) GetOperateCompliancePacksResult() *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult {
	return s.OperateCompliancePacksResult
}

func (s *DeleteAggregateCompliancePacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggregateCompliancePacksResponseBody) SetOperateCompliancePacksResult(v *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteAggregateCompliancePacksResponseBody {
	s.OperateCompliancePacksResult = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBody) SetRequestId(v string) *DeleteAggregateCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBody) Validate() error {
	if s.OperateCompliancePacksResult != nil {
		if err := s.OperateCompliancePacksResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult struct {
	OperateCompliancePacks []*DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks `json:"OperateCompliancePacks,omitempty" xml:"OperateCompliancePacks,omitempty" type:"Repeated"`
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) GetOperateCompliancePacks() []*DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	return s.OperateCompliancePacks
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) SetOperateCompliancePacks(v []*DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult {
	s.OperateCompliancePacks = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) Validate() error {
	if s.OperateCompliancePacks != nil {
		for _, item := range s.OperateCompliancePacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetCompliancePackId(v string) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetErrorCode(v string) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetSuccess(v bool) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.Success = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) Validate() error {
	return dara.Validate(s)
}
