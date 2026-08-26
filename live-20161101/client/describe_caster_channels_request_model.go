// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterChannelsRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterChannelsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterChannelsRequest
	GetRegionId() *string
}

type DescribeCasterChannelsRequest struct {
	// The ID of the production studio.
	//
	// - If you call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio, use the CasterId returned in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The name of a production studio in the list on the Cloud Production Studio page is the ID of the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCasterChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterChannelsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterChannelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterChannelsRequest) SetCasterId(v string) *DescribeCasterChannelsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterChannelsRequest) SetOwnerId(v int64) *DescribeCasterChannelsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterChannelsRequest) SetRegionId(v string) *DescribeCasterChannelsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterChannelsRequest) Validate() error {
	return dara.Validate(s)
}
