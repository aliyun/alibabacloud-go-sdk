// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *RestartCasterRequest
	GetCasterId() *string
	SetOwnerId(v int64) *RestartCasterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RestartCasterRequest
	GetRegionId() *string
}

type RestartCasterRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you created a production studio in the Live console, go to **Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
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

func (s RestartCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartCasterRequest) GoString() string {
	return s.String()
}

func (s *RestartCasterRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *RestartCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartCasterRequest) SetCasterId(v string) *RestartCasterRequest {
	s.CasterId = &v
	return s
}

func (s *RestartCasterRequest) SetOwnerId(v int64) *RestartCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartCasterRequest) SetRegionId(v string) *RestartCasterRequest {
	s.RegionId = &v
	return s
}

func (s *RestartCasterRequest) Validate() error {
	return dara.Validate(s)
}
