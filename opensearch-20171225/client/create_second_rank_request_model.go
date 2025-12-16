// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecondRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SecondRank) *CreateSecondRankRequest
	GetBody() *SecondRank
	SetDryRun(v bool) *CreateSecondRankRequest
	GetDryRun() *bool
}

type CreateSecondRankRequest struct {
	// The request body. For more information, see [SecondRank](https://help.aliyun.com/document_detail/170008.html).
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateSecondRankRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecondRankRequest) GoString() string {
	return s.String()
}

func (s *CreateSecondRankRequest) GetBody() *SecondRank {
	return s.Body
}

func (s *CreateSecondRankRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSecondRankRequest) SetBody(v *SecondRank) *CreateSecondRankRequest {
	s.Body = v
	return s
}

func (s *CreateSecondRankRequest) SetDryRun(v bool) *CreateSecondRankRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecondRankRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
