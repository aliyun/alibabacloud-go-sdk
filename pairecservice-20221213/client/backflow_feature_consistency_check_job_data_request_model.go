// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackflowFeatureConsistencyCheckJobDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckJobConfigId(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetFeatureConsistencyCheckJobConfigId() *string
	SetInstanceId(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetInstanceId() *string
	SetItemFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetItemFeatures() *string
	SetLogItemId(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetLogItemId() *string
	SetLogRequestId(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetLogRequestId() *string
	SetLogRequestTime(v int64) *BackflowFeatureConsistencyCheckJobDataRequest
	GetLogRequestTime() *int64
	SetLogUserId(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetLogUserId() *string
	SetSceneName(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetSceneName() *string
	SetScores(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetScores() *string
	SetServiceName(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetServiceName() *string
	SetUserFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest
	GetUserFeatures() *string
}

type BackflowFeatureConsistencyCheckJobDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
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
	// [\\"{\\\\\\"itemid\\\\\\":{\\\\\\"value\\\\\\":1010,\\\\\\"type\\\\\\":\\\\\\"string\\\\\\"}}\\"]
	ItemFeatures *string `json:"ItemFeatures,omitempty" xml:"ItemFeatures,omitempty"`
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
	//
	// example:
	//
	// video-feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"{\\\\\\"dbmtl_probs_is_valid_play\\\\\\":0.00032182207107543945,\\\\\\"dbmtl_y_play_time\\\\\\":0.0043269748210906982}\\"]
	Scores      *string `json:"Scores,omitempty" xml:"Scores,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"userid\\":{\\"value\\":1010,\\"type\\":\\"string\\"},\\"click_5_seq\\":{\\"value\\":\\"9001;9002;9003;9004;9005\\",\\"type\\":\\"string\\"}}
	UserFeatures *string `json:"UserFeatures,omitempty" xml:"UserFeatures,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataRequest) String() string {
	return dara.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataRequest) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetItemFeatures() *string {
	return s.ItemFeatures
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetLogItemId() *string {
	return s.LogItemId
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetLogRequestTime() *int64 {
	return s.LogRequestTime
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetLogUserId() *string {
	return s.LogUserId
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetScores() *string {
	return s.Scores
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) GetUserFeatures() *string {
	return s.UserFeatures
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetFeatureConsistencyCheckJobConfigId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetInstanceId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.InstanceId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetItemFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.ItemFeatures = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogItemId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogItemId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogRequestId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogRequestId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogRequestTime(v int64) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogRequestTime = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogUserId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogUserId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetSceneName(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.SceneName = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetScores(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.Scores = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetServiceName(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.ServiceName = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetUserFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.UserFeatures = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) Validate() error {
	return dara.Validate(s)
}
