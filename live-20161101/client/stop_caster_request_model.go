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
	// The ID of the production studio. Make sure that the specified CasterId is correct.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, obtain the ID from the CasterId parameter in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
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
