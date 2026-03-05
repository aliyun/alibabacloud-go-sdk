// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRecommendedCommoditiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppRecommendedCommoditiesRequest
	GetBizId() *string
	SetScene(v string) *GetAppRecommendedCommoditiesRequest
	GetScene() *string
}

type GetAppRecommendedCommoditiesRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// DataworksManualTask
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s GetAppRecommendedCommoditiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesRequest) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppRecommendedCommoditiesRequest) GetScene() *string {
	return s.Scene
}

func (s *GetAppRecommendedCommoditiesRequest) SetBizId(v string) *GetAppRecommendedCommoditiesRequest {
	s.BizId = &v
	return s
}

func (s *GetAppRecommendedCommoditiesRequest) SetScene(v string) *GetAppRecommendedCommoditiesRequest {
	s.Scene = &v
	return s
}

func (s *GetAppRecommendedCommoditiesRequest) Validate() error {
	return dara.Validate(s)
}
