// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultimodeCmsUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMultimodeCmsUrlRequest
	GetClusterId() *string
	SetRegionId(v string) *GetMultimodeCmsUrlRequest
	GetRegionId() *string
}

type GetMultimodeCmsUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetMultimodeCmsUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultimodeCmsUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMultimodeCmsUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMultimodeCmsUrlRequest) SetClusterId(v string) *GetMultimodeCmsUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetRegionId(v string) *GetMultimodeCmsUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) Validate() error {
	return dara.Validate(s)
}
