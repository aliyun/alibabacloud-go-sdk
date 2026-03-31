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
	// The results of the delete operations.
	OperateCompliancePacksResult *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult `json:"OperateCompliancePacksResult,omitempty" xml:"OperateCompliancePacksResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// An array that contains the deleted compliance packages.
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
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The error code returned.
	//
	// 	- If the compliance package is deleted, no error code is returned.
	//
	// 	- If the compliance package fails to be deleted, an error code is returned. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// CompliancePackAlreadyPending
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
