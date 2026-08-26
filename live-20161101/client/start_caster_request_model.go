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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter in the response of CreateCaster.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
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
