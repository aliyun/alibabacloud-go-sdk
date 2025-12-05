// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneBaseLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaseline(v *GetPtsSceneBaseLineResponseBodyBaseline) *GetPtsSceneBaseLineResponseBody
	GetBaseline() *GetPtsSceneBaseLineResponseBodyBaseline
	SetCode(v string) *GetPtsSceneBaseLineResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPtsSceneBaseLineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsSceneBaseLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPtsSceneBaseLineResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetPtsSceneBaseLineResponseBody
	GetSceneId() *string
	SetSuccess(v bool) *GetPtsSceneBaseLineResponseBody
	GetSuccess() *bool
}

type GetPtsSceneBaseLineResponseBody struct {
	// Baseline data
	Baseline *GetPtsSceneBaseLineResponseBodyBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F7D2CE0-AE4C-4143-955A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// NHG67BF
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false:
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBody) GetBaseline() *GetPtsSceneBaseLineResponseBodyBaseline {
	return s.Baseline
}

func (s *GetPtsSceneBaseLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsSceneBaseLineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsSceneBaseLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsSceneBaseLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsSceneBaseLineResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneBaseLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsSceneBaseLineResponseBody) SetBaseline(v *GetPtsSceneBaseLineResponseBodyBaseline) *GetPtsSceneBaseLineResponseBody {
	s.Baseline = v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetCode(v string) *GetPtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetMessage(v string) *GetPtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetRequestId(v string) *GetPtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetSceneId(v string) *GetPtsSceneBaseLineResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetSuccess(v bool) *GetPtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) Validate() error {
	if s.Baseline != nil {
		if err := s.Baseline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPtsSceneBaseLineResponseBodyBaseline struct {
	// null
	ApiBaselines []*GetPtsSceneBaseLineResponseBodyBaselineApiBaselines `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty" type:"Repeated"`
	// Scenario
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// null
	SceneBaseline *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty" type:"Struct"`
}

func (s GetPtsSceneBaseLineResponseBodyBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaseline) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) GetApiBaselines() []*GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	return s.ApiBaselines
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) GetName() *string {
	return s.Name
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) GetSceneBaseline() *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	return s.SceneBaseline
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetApiBaselines(v []*GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.ApiBaselines = v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetName(v string) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.Name = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetSceneBaseline(v *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.SceneBaseline = v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) Validate() error {
	if s.ApiBaselines != nil {
		for _, item := range s.ApiBaselines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneBaseline != nil {
		if err := s.SceneBaseline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPtsSceneBaseLineResponseBodyBaselineApiBaselines struct {
	// Average RT
	//
	// example:
	//
	// 10
	AvgRt *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// null
	//
	// example:
	//
	// 1000
	AvgTps *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// null
	//
	// example:
	//
	// 100
	FailCountBiz *int64 `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	// Failures
	//
	// example:
	//
	// 100
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// The API ID.
	//
	// example:
	//
	// 76543
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// null
	//
	// example:
	//
	// 50
	MaxRt *int32 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// null
	//
	// example:
	//
	// 8
	MinRt *int32 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	// The name of the API operation.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// null
	//
	// example:
	//
	// 40
	Seg90Rt *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// null
	//
	// example:
	//
	// 50
	Seg99Rt *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// null
	//
	// example:
	//
	// 0.1
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	// request success rate
	//
	// example:
	//
	// 0.9
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetAvgRt() *float32 {
	return s.AvgRt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetAvgTps() *float32 {
	return s.AvgTps
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetFailCountBiz() *int64 {
	return s.FailCountBiz
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetId() *int64 {
	return s.Id
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetMaxRt() *int32 {
	return s.MaxRt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetMinRt() *int32 {
	return s.MinRt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetName() *string {
	return s.Name
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetSeg90Rt() *float32 {
	return s.Seg90Rt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetSeg99Rt() *float32 {
	return s.Seg99Rt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetSuccessRateBiz() *float32 {
	return s.SuccessRateBiz
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GetSuccessRateReq() *float32 {
	return s.SuccessRateReq
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetAvgRt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.AvgRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetAvgTps(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.AvgTps = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetFailCountBiz(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetFailCountReq(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetId(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Id = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetMaxRt(v int32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.MaxRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetMinRt(v int32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.MinRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetName(v string) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Name = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSeg90Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSeg99Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSuccessRateBiz(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSuccessRateReq(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.SuccessRateReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline struct {
	// Average RT
	//
	// example:
	//
	// 10
	AvgRt *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// null
	//
	// example:
	//
	// 1000
	AvgTps *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// null
	//
	// example:
	//
	// 1000
	FailCountBiz *int64 `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	// Failures
	//
	// example:
	//
	// 1000
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// null
	//
	// example:
	//
	// 10
	Seg90Rt *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// null
	//
	// example:
	//
	// 10
	Seg99Rt *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// null
	//
	// example:
	//
	// 0.1
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	// request success rate
	//
	// example:
	//
	// 0.9
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetAvgRt() *float32 {
	return s.AvgRt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetAvgTps() *float32 {
	return s.AvgTps
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetFailCountBiz() *int64 {
	return s.FailCountBiz
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetSeg90Rt() *float32 {
	return s.Seg90Rt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetSeg99Rt() *float32 {
	return s.Seg99Rt
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetSuccessRateBiz() *float32 {
	return s.SuccessRateBiz
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GetSuccessRateReq() *float32 {
	return s.SuccessRateReq
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetAvgRt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.AvgRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetAvgTps(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.AvgTps = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetFailCountBiz(v int64) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetFailCountReq(v int64) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSeg90Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSeg99Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSuccessRateBiz(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSuccessRateReq(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.SuccessRateReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) Validate() error {
	return dara.Validate(s)
}
