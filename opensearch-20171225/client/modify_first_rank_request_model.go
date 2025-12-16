// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirstRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *FirstRank) *ModifyFirstRankRequest
	GetBody() *FirstRank
	SetDryRun(v bool) *ModifyFirstRankRequest
	GetDryRun() *bool
}

type ModifyFirstRankRequest struct {
	// The request body.
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyFirstRankRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirstRankRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankRequest) GetBody() *FirstRank {
	return s.Body
}

func (s *ModifyFirstRankRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyFirstRankRequest) SetBody(v *FirstRank) *ModifyFirstRankRequest {
	s.Body = v
	return s
}

func (s *ModifyFirstRankRequest) SetDryRun(v bool) *ModifyFirstRankRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFirstRankRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
