// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *StartCasterRequest
	GetCasterId() *string
	SetOwnerId(v int64) *StartCasterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartCasterRequest
	GetRegionId() *string
}

type StartCasterRequest struct {
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
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCasterRequest) GoString() string {
	return s.String()
}

func (s *StartCasterRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *StartCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartCasterRequest) SetCasterId(v string) *StartCasterRequest {
	s.CasterId = &v
	return s
}

func (s *StartCasterRequest) SetOwnerId(v int64) *StartCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *StartCasterRequest) SetRegionId(v string) *StartCasterRequest {
	s.RegionId = &v
	return s
}

func (s *StartCasterRequest) Validate() error {
	return dara.Validate(s)
}
