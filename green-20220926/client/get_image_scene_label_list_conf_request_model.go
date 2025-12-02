// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelListConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageServiceCode(v string) *GetImageSceneLabelListConfRequest
	GetImageServiceCode() *string
	SetRegionId(v string) *GetImageSceneLabelListConfRequest
	GetRegionId() *string
}

type GetImageSceneLabelListConfRequest struct {
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ImageServiceCode *string `json:"ImageServiceCode,omitempty" xml:"ImageServiceCode,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetImageSceneLabelListConfRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelListConfRequest) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfRequest) GetImageServiceCode() *string {
	return s.ImageServiceCode
}

func (s *GetImageSceneLabelListConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageSceneLabelListConfRequest) SetImageServiceCode(v string) *GetImageSceneLabelListConfRequest {
	s.ImageServiceCode = &v
	return s
}

func (s *GetImageSceneLabelListConfRequest) SetRegionId(v string) *GetImageSceneLabelListConfRequest {
	s.RegionId = &v
	return s
}

func (s *GetImageSceneLabelListConfRequest) Validate() error {
	return dara.Validate(s)
}
