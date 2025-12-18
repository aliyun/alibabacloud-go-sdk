// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompliancePacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateCompliancePacksResult(v *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteCompliancePacksResponseBody
	GetOperateCompliancePacksResult() *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult
	SetRequestId(v string) *DeleteCompliancePacksResponseBody
	GetRequestId() *string
}

type DeleteCompliancePacksResponseBody struct {
	// The returned result.
	OperateCompliancePacksResult *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult `json:"OperateCompliancePacksResult,omitempty" xml:"OperateCompliancePacksResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCompliancePacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBody) GetOperateCompliancePacksResult() *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult {
	return s.OperateCompliancePacksResult
}

func (s *DeleteCompliancePacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCompliancePacksResponseBody) SetOperateCompliancePacksResult(v *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteCompliancePacksResponseBody {
	s.OperateCompliancePacksResult = v
	return s
}

func (s *DeleteCompliancePacksResponseBody) SetRequestId(v string) *DeleteCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCompliancePacksResponseBody) Validate() error {
	if s.OperateCompliancePacksResult != nil {
		if err := s.OperateCompliancePacksResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCompliancePacksResponseBodyOperateCompliancePacksResult struct {
	// An array that contains compliance packages that are deleted.
	OperateCompliancePacks []*DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks `json:"OperateCompliancePacks,omitempty" xml:"OperateCompliancePacks,omitempty" type:"Repeated"`
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) GetOperateCompliancePacks() []*DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	return s.OperateCompliancePacks
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) SetOperateCompliancePacks(v []*DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult {
	s.OperateCompliancePacks = v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) Validate() error {
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

type DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetCompliancePackId(v string) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetErrorCode(v string) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetSuccess(v bool) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.Success = &v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) Validate() error {
	return dara.Validate(s)
}
