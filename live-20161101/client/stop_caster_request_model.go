// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *StopCasterRequest
	GetCasterId() *string
	SetOwnerId(v int64) *StopCasterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopCasterRequest
	GetRegionId() *string
}

type StopCasterRequest struct {
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

func (s StopCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCasterRequest) GoString() string {
	return s.String()
}

func (s *StopCasterRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *StopCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopCasterRequest) SetCasterId(v string) *StopCasterRequest {
	s.CasterId = &v
	return s
}

func (s *StopCasterRequest) SetOwnerId(v int64) *StopCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCasterRequest) SetRegionId(v string) *StopCasterRequest {
	s.RegionId = &v
	return s
}

func (s *StopCasterRequest) Validate() error {
	return dara.Validate(s)
}
