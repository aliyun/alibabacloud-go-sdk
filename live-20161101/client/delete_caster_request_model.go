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
