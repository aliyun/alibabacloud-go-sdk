// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetImageSceneLabelConfRequest
	GetRegionId() *string
}

type GetImageSceneLabelConfRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetImageSceneLabelConfRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelConfRequest) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageSceneLabelConfRequest) SetRegionId(v string) *GetImageSceneLabelConfRequest {
	s.RegionId = &v
	return s
}

func (s *GetImageSceneLabelConfRequest) Validate() error {
	return dara.Validate(s)
}
