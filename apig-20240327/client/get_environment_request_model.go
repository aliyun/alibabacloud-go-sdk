// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWithStatistics(v bool) *GetEnvironmentRequest
	GetWithStatistics() *bool
	SetWithVpcInfo(v bool) *GetEnvironmentRequest
	GetWithVpcInfo() *bool
}

type GetEnvironmentRequest struct {
	// Indicates whether to return online resource info.
	//
	// example:
	//
	// true
	WithStatistics *bool `json:"withStatistics,omitempty" xml:"withStatistics,omitempty"`
	// Option for vpc info.
	//
	// example:
	//
	// true
	WithVpcInfo *bool `json:"withVpcInfo,omitempty" xml:"withVpcInfo,omitempty"`
}

func (s GetEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *GetEnvironmentRequest) GetWithStatistics() *bool {
	return s.WithStatistics
}

func (s *GetEnvironmentRequest) GetWithVpcInfo() *bool {
	return s.WithVpcInfo
}

func (s *GetEnvironmentRequest) SetWithStatistics(v bool) *GetEnvironmentRequest {
	s.WithStatistics = &v
	return s
}

func (s *GetEnvironmentRequest) SetWithVpcInfo(v bool) *GetEnvironmentRequest {
	s.WithVpcInfo = &v
	return s
}

func (s *GetEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
