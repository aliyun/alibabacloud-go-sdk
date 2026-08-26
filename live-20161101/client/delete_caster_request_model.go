// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DeleteCasterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterRequest
	GetRegionId() *string
}

type DeleteCasterRequest struct {
	// The ID of the production studio. Make sure that you specify the correct CasterId.
	//
	// - If you created a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, obtain the CasterId from the response.
	//
	// - If you created a production studio in the ApsaraVideo Live console, go to the **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > On the Cloud Production Studio page, the name of a production studio in the list is its ID.
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

func (s DeleteCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterRequest) SetCasterId(v string) *DeleteCasterRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterRequest) SetOwnerId(v int64) *DeleteCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterRequest) SetRegionId(v string) *DeleteCasterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterRequest) Validate() error {
	return dara.Validate(s)
}
