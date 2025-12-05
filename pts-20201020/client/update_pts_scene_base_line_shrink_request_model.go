// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePtsSceneBaseLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiBaselinesShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest
	GetApiBaselinesShrink() *string
	SetSceneBaselineShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest
	GetSceneBaselineShrink() *string
	SetSceneId(v string) *UpdatePtsSceneBaseLineShrinkRequest
	GetSceneId() *string
}

type UpdatePtsSceneBaseLineShrinkRequest struct {
	// null null
	//
	// example:
	//
	// [{"avgRt":1,"avgTps":1,"failCountBiz":1,"failCountReq":182381,"id":362447,"maxRt":3051,"minRt":0,"name":"1-1","seg50Rt":1,"seg75Rt":1,"seg90Rt":1,"seg99Rt":3,"successRateBiz":1,"successRateReq":0,"timingConnAvg":0},{"avgRt":1.06,"avgTps":1,"failCountBiz":0,"failCountReq":151143,"id":362446,"maxRt":3068,"minRt":0,"name":"dd","seg50Rt":1,"seg75Rt":1,"seg90Rt":1,"seg99Rt":2,"successRateBiz":1,"successRateReq":0}]
	ApiBaselinesShrink *string `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty"`
	// null null
	//
	// example:
	//
	// {"avgRt":1,"avgTps":1,"failCountBiz":1,"failCountReq":1,"seg90Rt":1,"seg99Rt":2,"successRateBiz":0.5,"successRateReq":1}
	SceneBaselineShrink *string `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty"`
	// The ID of the scene. For more information, see the [table](https://help.aliyun.com/document_detail/201321.html) provided in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// NB54CV
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdatePtsSceneBaseLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePtsSceneBaseLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) GetApiBaselinesShrink() *string {
	return s.ApiBaselinesShrink
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) GetSceneBaselineShrink() *string {
	return s.SceneBaselineShrink
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetApiBaselinesShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.ApiBaselinesShrink = &v
	return s
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetSceneBaselineShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.SceneBaselineShrink = &v
	return s
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetSceneId(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
