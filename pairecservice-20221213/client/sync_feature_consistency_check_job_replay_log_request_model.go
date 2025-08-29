// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFeatureConsistencyCheckJobReplayLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetContextFeatures() *string
	SetFeatureConsistencyCheckJobConfigId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetFeatureConsistencyCheckJobConfigId() *string
	SetGeneratedFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetGeneratedFeatures() *string
	SetInstanceId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetInstanceId() *string
	SetLogItemId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetLogItemId() *string
	SetLogRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetLogRequestId() *string
	SetLogRequestTime(v int64) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetLogRequestTime() *int64
	SetLogUserId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetLogUserId() *string
	SetRawFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetRawFeatures() *string
	SetSceneName(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest
	GetSceneName() *string
}

type SyncFeatureConsistencyCheckJobReplayLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{\\"Value\\":{\\"FloatFeature\\":0.1}}]
	ContextFeatures *string `json:"ContextFeatures,omitempty" xml:"ContextFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// week_day:1 | userid:3 | itemid:9001 | cate:cat1 | click_5_seq__cate:cat1
	GeneratedFeatures *string `json:"GeneratedFeatures,omitempty" xml:"GeneratedFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1693900981465
	LogRequestTime *int64 `json:"LogRequestTime,omitempty" xml:"LogRequestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// This parameter is required.
	RawFeatures *string `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video-feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogRequest) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetContextFeatures() *string {
	return s.ContextFeatures
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetGeneratedFeatures() *string {
	return s.GeneratedFeatures
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetLogItemId() *string {
	return s.LogItemId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetLogRequestTime() *int64 {
	return s.LogRequestTime
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetLogUserId() *string {
	return s.LogUserId
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetRawFeatures() *string {
	return s.RawFeatures
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetContextFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.ContextFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetFeatureConsistencyCheckJobConfigId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetGeneratedFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.GeneratedFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetInstanceId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogItemId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogItemId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogRequestId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogRequestTime(v int64) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogRequestTime = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogUserId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogUserId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetRawFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.RawFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetSceneName(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.SceneName = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) Validate() error {
	return dara.Validate(s)
}
