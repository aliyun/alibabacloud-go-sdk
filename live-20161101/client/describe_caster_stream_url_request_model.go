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
	// The ID of the production studio. Make sure that the CasterId is correct.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, obtain the ID from the CasterId parameter in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **ApsaraVideo Live console*	- and choose **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The name of a production studio in the list on the Cloud Production Studio page is its ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
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
