// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DescribeVerifyResultRequest
	GetBizId() *string
	SetBizType(v string) *DescribeVerifyResultRequest
	GetBizType() *string
}

type DescribeVerifyResultRequest struct {
	// The verification ID. A unique ID that identifies a verification task. The value can be up to 64 characters in length. For a single verification task, the system supports unlimited submissions until the verification is passed and the task is completed.
	//
	// > Use a different BizId for each verification task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125b****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The business scenario identifier of the ID Verification service.
	//
	// This parameter is required.
	//
	// example:
	//
	// FVBioOnlyTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s DescribeVerifyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultRequest) GetBizId() *string {
	return s.BizId
}

func (s *DescribeVerifyResultRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeVerifyResultRequest) SetBizId(v string) *DescribeVerifyResultRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyResultRequest) SetBizType(v string) *DescribeVerifyResultRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyResultRequest) Validate() error {
	return dara.Validate(s)
}
