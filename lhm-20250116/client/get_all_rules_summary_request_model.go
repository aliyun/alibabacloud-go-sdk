// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllRulesSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetAllRulesSummaryRequest
	GetSource() *string
	SetTarget(v string) *GetAllRulesSummaryRequest
	GetTarget() *string
}

type GetAllRulesSummaryRequest struct {
	// The source dialect.
	//
	// example:
	//
	// postgresql
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target dialect.
	//
	// example:
	//
	// hologres
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s GetAllRulesSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllRulesSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetAllRulesSummaryRequest) GetSource() *string {
	return s.Source
}

func (s *GetAllRulesSummaryRequest) GetTarget() *string {
	return s.Target
}

func (s *GetAllRulesSummaryRequest) SetSource(v string) *GetAllRulesSummaryRequest {
	s.Source = &v
	return s
}

func (s *GetAllRulesSummaryRequest) SetTarget(v string) *GetAllRulesSummaryRequest {
	s.Target = &v
	return s
}

func (s *GetAllRulesSummaryRequest) Validate() error {
	return dara.Validate(s)
}
