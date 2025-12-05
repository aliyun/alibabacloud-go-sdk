// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePtsSceneBaseLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiBaselines(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest
	GetApiBaselines() map[string]interface{}
	SetSceneBaseline(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest
	GetSceneBaseline() map[string]interface{}
	SetSceneId(v string) *UpdatePtsSceneBaseLineRequest
	GetSceneId() *string
}

type UpdatePtsSceneBaseLineRequest struct {
	// null null
	//
	// example:
	//
	// [{"avgRt":1,"avgTps":1,"failCountBiz":1,"failCountReq":182381,"id":362447,"maxRt":3051,"minRt":0,"name":"1-1","seg50Rt":1,"seg75Rt":1,"seg90Rt":1,"seg99Rt":3,"successRateBiz":1,"successRateReq":0,"timingConnAvg":0},{"avgRt":1.06,"avgTps":1,"failCountBiz":0,"failCountReq":151143,"id":362446,"maxRt":3068,"minRt":0,"name":"dd","seg50Rt":1,"seg75Rt":1,"seg90Rt":1,"seg99Rt":2,"successRateBiz":1,"successRateReq":0}]
	ApiBaselines map[string]interface{} `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty"`
	// null null
	//
	// example:
	//
	// {"avgRt":1,"avgTps":1,"failCountBiz":1,"failCountReq":1,"seg90Rt":1,"seg99Rt":2,"successRateBiz":0.5,"successRateReq":1}
	SceneBaseline map[string]interface{} `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty"`
	// The ID of the scene. For more information, see the [table](https://help.aliyun.com/document_detail/201321.html) provided in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// NB54CV
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdatePtsSceneBaseLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineRequest) GetApiBaselines() map[string]interface{} {
	return s.ApiBaselines
}

func (s *UpdatePtsSceneBaseLineRequest) GetSceneBaseline() map[string]interface{} {
	return s.SceneBaseline
}

func (s *UpdatePtsSceneBaseLineRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdatePtsSceneBaseLineRequest) SetApiBaselines(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest {
	s.ApiBaselines = v
	return s
}

func (s *UpdatePtsSceneBaseLineRequest) SetSceneBaseline(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest {
	s.SceneBaseline = v
	return s
}

func (s *UpdatePtsSceneBaseLineRequest) SetSceneId(v string) *UpdatePtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

func (s *UpdatePtsSceneBaseLineRequest) Validate() error {
	return dara.Validate(s)
}
