// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirstRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *FirstRank) *CreateFirstRankRequest
	GetBody() *FirstRank
	SetDryRun(v bool) *CreateFirstRankRequest
	GetDryRun() *bool
}

type CreateFirstRankRequest struct {
	// The request body that contains the parameters of the rough sort expression.
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateFirstRankRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirstRankRequest) GoString() string {
	return s.String()
}

func (s *CreateFirstRankRequest) GetBody() *FirstRank {
	return s.Body
}

func (s *CreateFirstRankRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFirstRankRequest) SetBody(v *FirstRank) *CreateFirstRankRequest {
	s.Body = v
	return s
}

func (s *CreateFirstRankRequest) SetDryRun(v bool) *CreateFirstRankRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFirstRankRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
