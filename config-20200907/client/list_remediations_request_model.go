// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleIds(v string) *ListRemediationsRequest
	GetConfigRuleIds() *string
	SetPageNumber(v int64) *ListRemediationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationsRequest
	GetPageSize() *int64
}

type ListRemediationsRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationsRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *ListRemediationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationsRequest) SetConfigRuleIds(v string) *ListRemediationsRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *ListRemediationsRequest) SetPageNumber(v int64) *ListRemediationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationsRequest) SetPageSize(v int64) *ListRemediationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
