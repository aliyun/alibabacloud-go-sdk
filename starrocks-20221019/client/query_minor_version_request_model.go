// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *QueryMinorVersionRequest
	GetRegionId() *string
	SetVersion(v string) *QueryMinorVersionRequest
	GetVersion() *string
}

type QueryMinorVersionRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *QueryMinorVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryMinorVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *QueryMinorVersionRequest) SetRegionId(v string) *QueryMinorVersionRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMinorVersionRequest) SetVersion(v string) *QueryMinorVersionRequest {
	s.Version = &v
	return s
}

func (s *QueryMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
