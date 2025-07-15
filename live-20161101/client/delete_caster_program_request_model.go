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
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
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
