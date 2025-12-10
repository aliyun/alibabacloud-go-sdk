// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetAttachmentsForControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyRequest
	GetPageSize() *int32
	SetPolicyId(v string) *ListTargetAttachmentsForControlPolicyRequest
	GetPolicyId() *string
}

type ListTargetAttachmentsForControlPolicyRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTargetAttachmentsForControlPolicyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTargetAttachmentsForControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPolicyId(v string) *ListTargetAttachmentsForControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
