// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetManagedRuleIdentifier(v string) *ListRemediationTemplatesRequest
	GetManagedRuleIdentifier() *string
	SetPageNumber(v int64) *ListRemediationTemplatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationTemplatesRequest
	GetPageSize() *int64
	SetRemediationType(v string) *ListRemediationTemplatesRequest
	GetRemediationType() *string
}

type ListRemediationTemplatesRequest struct {
	// if can be null:
	// true
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RemediationType       *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListRemediationTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesRequest) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *ListRemediationTemplatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationTemplatesRequest) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListRemediationTemplatesRequest) SetManagedRuleIdentifier(v string) *ListRemediationTemplatesRequest {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *ListRemediationTemplatesRequest) SetPageNumber(v int64) *ListRemediationTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationTemplatesRequest) SetPageSize(v int64) *ListRemediationTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRemediationTemplatesRequest) SetRemediationType(v string) *ListRemediationTemplatesRequest {
	s.RemediationType = &v
	return s
}

func (s *ListRemediationTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
