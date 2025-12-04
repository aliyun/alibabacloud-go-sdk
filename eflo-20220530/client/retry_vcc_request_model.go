// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *RetryVccRequest
	GetRegionId() *string
	SetVccId(v string) *RetryVccRequest
	GetVccId() *string
}

type RetryVccRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s RetryVccRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryVccRequest) GoString() string {
	return s.String()
}

func (s *RetryVccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryVccRequest) GetVccId() *string {
	return s.VccId
}

func (s *RetryVccRequest) SetRegionId(v string) *RetryVccRequest {
	s.RegionId = &v
	return s
}

func (s *RetryVccRequest) SetVccId(v string) *RetryVccRequest {
	s.VccId = &v
	return s
}

func (s *RetryVccRequest) Validate() error {
	return dara.Validate(s)
}
