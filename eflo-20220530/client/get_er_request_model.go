// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *GetErRequest
	GetErId() *string
	SetRegionId(v string) *GetErRequest
	GetRegionId() *string
}

type GetErRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErRequest) GoString() string {
	return s.String()
}

func (s *GetErRequest) GetErId() *string {
	return s.ErId
}

func (s *GetErRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErRequest) SetErId(v string) *GetErRequest {
	s.ErId = &v
	return s
}

func (s *GetErRequest) SetRegionId(v string) *GetErRequest {
	s.RegionId = &v
	return s
}

func (s *GetErRequest) Validate() error {
	return dara.Validate(s)
}
