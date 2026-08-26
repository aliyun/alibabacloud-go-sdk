// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterProgramRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DeleteCasterProgramRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterProgramRequest
	GetRegionId() *string
}

type DeleteCasterProgramRequest struct {
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
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

func (s DeleteCasterProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterProgramRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterProgramRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterProgramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterProgramRequest) SetCasterId(v string) *DeleteCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterProgramRequest) SetOwnerId(v int64) *DeleteCasterProgramRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterProgramRequest) SetRegionId(v string) *DeleteCasterProgramRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterProgramRequest) Validate() error {
	return dara.Validate(s)
}
