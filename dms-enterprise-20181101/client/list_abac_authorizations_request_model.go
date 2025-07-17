// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAbacAuthorizationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAbacAuthorizationsRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListAbacAuthorizationsRequest
	GetPolicyId() *string
	SetPolicySource(v string) *ListAbacAuthorizationsRequest
	GetPolicySource() *string
	SetTid(v int64) *ListAbacAuthorizationsRequest
	GetTid() *int64
}

type ListAbacAuthorizationsRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// USER_DEFINE
	PolicySource *string `json:"PolicySource,omitempty" xml:"PolicySource,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAbacAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAbacAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *ListAbacAuthorizationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAbacAuthorizationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAbacAuthorizationsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListAbacAuthorizationsRequest) GetPolicySource() *string {
	return s.PolicySource
}

func (s *ListAbacAuthorizationsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAbacAuthorizationsRequest) SetPageNumber(v int64) *ListAbacAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPageSize(v int64) *ListAbacAuthorizationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPolicyId(v string) *ListAbacAuthorizationsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetPolicySource(v string) *ListAbacAuthorizationsRequest {
	s.PolicySource = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) SetTid(v int64) *ListAbacAuthorizationsRequest {
	s.Tid = &v
	return s
}

func (s *ListAbacAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
