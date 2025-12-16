// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecondRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SecondRank) *ModifySecondRankRequest
	GetBody() *SecondRank
	SetDryRun(v bool) *ModifySecondRankRequest
	GetDryRun() *bool
}

type ModifySecondRankRequest struct {
	// The request parameters.
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifySecondRankRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecondRankRequest) GoString() string {
	return s.String()
}

func (s *ModifySecondRankRequest) GetBody() *SecondRank {
	return s.Body
}

func (s *ModifySecondRankRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifySecondRankRequest) SetBody(v *SecondRank) *ModifySecondRankRequest {
	s.Body = v
	return s
}

func (s *ModifySecondRankRequest) SetDryRun(v bool) *ModifySecondRankRequest {
	s.DryRun = &v
	return s
}

func (s *ModifySecondRankRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
