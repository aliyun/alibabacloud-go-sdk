// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWithStatistics(v bool) *GetDomainRequest
	GetWithStatistics() *bool
}

type GetDomainRequest struct {
	// Specifies whether to return online resource information.
	//
	// example:
	//
	// true
	WithStatistics *bool `json:"withStatistics,omitempty" xml:"withStatistics,omitempty"`
}

func (s GetDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) GetWithStatistics() *bool {
	return s.WithStatistics
}

func (s *GetDomainRequest) SetWithStatistics(v bool) *GetDomainRequest {
	s.WithStatistics = &v
	return s
}

func (s *GetDomainRequest) Validate() error {
	return dara.Validate(s)
}
