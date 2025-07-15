// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterStreamUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterStreamUrlRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterStreamUrlRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterStreamUrlRequest
	GetRegionId() *string
}

type DescribeCasterStreamUrlRequest struct {
	// The ID of the production studio.
	//
	// If you create a production studio through the [CreateCaster](~~69338#doc-api-live-CreateCaster~~) interface, check the value of the CasterId parameter in the response.
	//
	// If you create a production studio through the ApsaraVideo Live Console, log in to the console, then check the ID of the production studio through the following path:
	//
	// Production Studios > Production Studio Management
	//
	// >  The CasterId is reflected in the Name column on the Production Studio Management page.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCasterStreamUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterStreamUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterStreamUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterStreamUrlRequest) SetCasterId(v string) *DescribeCasterStreamUrlRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterStreamUrlRequest) SetOwnerId(v int64) *DescribeCasterStreamUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterStreamUrlRequest) SetRegionId(v string) *DescribeCasterStreamUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterStreamUrlRequest) Validate() error {
	return dara.Validate(s)
}
