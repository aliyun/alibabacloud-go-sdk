// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterConfigRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterConfigRequest
	GetRegionId() *string
}

type DescribeCasterConfigRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The name of the production studio in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCasterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterConfigRequest) SetCasterId(v string) *DescribeCasterConfigRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterConfigRequest) SetOwnerId(v int64) *DescribeCasterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterConfigRequest) SetRegionId(v string) *DescribeCasterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterConfigRequest) Validate() error {
	return dara.Validate(s)
}
