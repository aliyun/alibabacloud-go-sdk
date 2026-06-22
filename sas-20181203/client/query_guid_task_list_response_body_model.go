// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGuidTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGuideTaskConfigList(v []*QueryGuidTaskListResponseBodyGuideTaskConfigList) *QueryGuidTaskListResponseBody
	GetGuideTaskConfigList() []*QueryGuidTaskListResponseBodyGuideTaskConfigList
	SetRequestId(v string) *QueryGuidTaskListResponseBody
	GetRequestId() *string
}

type QueryGuidTaskListResponseBody struct {
	// The list of beginner task information.
	GuideTaskConfigList []*QueryGuidTaskListResponseBodyGuideTaskConfigList `json:"GuideTaskConfigList,omitempty" xml:"GuideTaskConfigList,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGuidTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGuidTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGuidTaskListResponseBody) GetGuideTaskConfigList() []*QueryGuidTaskListResponseBodyGuideTaskConfigList {
	return s.GuideTaskConfigList
}

func (s *QueryGuidTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGuidTaskListResponseBody) SetGuideTaskConfigList(v []*QueryGuidTaskListResponseBodyGuideTaskConfigList) *QueryGuidTaskListResponseBody {
	s.GuideTaskConfigList = v
	return s
}

func (s *QueryGuidTaskListResponseBody) SetRequestId(v string) *QueryGuidTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGuidTaskListResponseBody) Validate() error {
	if s.GuideTaskConfigList != nil {
		for _, item := range s.GuideTaskConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryGuidTaskListResponseBodyGuideTaskConfigList struct {
	// The reward information for task completion.
	RewardData *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData `json:"RewardData,omitempty" xml:"RewardData,omitempty" type:"Struct"`
	// The security score increase that can be gained by completing this task.
	//
	// example:
	//
	// 80
	SecurityScore *int32 `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
	// The task status. Valid values:
	//
	// - **0**: Closed.
	//
	// - **1**: In progress.
	//
	// - **2**: Completed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-000d8slfgx4p40kb64ad
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The node name. Valid values:
	//
	// - **guid_task_security_score_promote_video**: the node of watching the beginner quick start video
	//
	// - **guide_sub_task_config_defence_hbr**: the anti-ransomware configuration node for servers
	//
	// - **guide_sub_task_config_uni_defence_hbr**: the anti-ransomware configuration node for databases
	//
	// - **guid_task_log_analysis_config**: the log analysis node
	//
	// - **guide_sub_task_web_lock_config**: the web tamper-proofing node
	//
	// - **guide_sub_task_config_anti_crack**: the anti-brute-force attacks node
	//
	// - **guid_task_container_security_video**: the container security video node
	//
	// - **guid_task_container_image_scan_config**: the container image scan node
	//
	// - **guid_task_k8s_log_analysis_config**: the Kubernetes threat detection node
	//
	// - **guid_task_container_network**: the container visualization node
	//
	// - **guide_sub_task_config_add_collection**: the node of adding the console to favorites
	//
	// - **guide_sub_task_vul_scan**: the vulnerability scanning node
	//
	// - **guide_sub_task_virusKill**: the virus scan node.
	//
	// example:
	//
	// guide_sub_task_config_add_collection
	TaskTypeName *string `json:"TaskTypeName,omitempty" xml:"TaskTypeName,omitempty"`
}

func (s QueryGuidTaskListResponseBodyGuideTaskConfigList) String() string {
	return dara.Prettify(s)
}

func (s QueryGuidTaskListResponseBodyGuideTaskConfigList) GoString() string {
	return s.String()
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) GetRewardData() *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData {
	return s.RewardData
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) GetSecurityScore() *int32 {
	return s.SecurityScore
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) GetTaskId() *int32 {
	return s.TaskId
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) GetTaskTypeName() *string {
	return s.TaskTypeName
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) SetRewardData(v *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) *QueryGuidTaskListResponseBodyGuideTaskConfigList {
	s.RewardData = v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) SetSecurityScore(v int32) *QueryGuidTaskListResponseBodyGuideTaskConfigList {
	s.SecurityScore = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) SetStatus(v int32) *QueryGuidTaskListResponseBodyGuideTaskConfigList {
	s.Status = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) SetTaskId(v int32) *QueryGuidTaskListResponseBodyGuideTaskConfigList {
	s.TaskId = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) SetTaskTypeName(v string) *QueryGuidTaskListResponseBodyGuideTaskConfigList {
	s.TaskTypeName = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigList) Validate() error {
	if s.RewardData != nil {
		if err := s.RewardData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData struct {
	// The claim status. Valid values:
	//
	// - **1**: Not claimed.
	//
	// - **2**: Claimed.
	//
	// example:
	//
	// 1
	IsRewardTaked *string `json:"IsRewardTaked,omitempty" xml:"IsRewardTaked,omitempty"`
	// The reward name. Valid values:
	//
	// - **addTrialDay**: trial days reward
	//
	// - **addAntiRansomwareCapacity**: anti-ransomware capacity reward
	//
	// - **addImageScanAuthCount**: image scan authorization quota reward
	//
	// - **addWebLockAuthCount**: web tamper-proofing authorization quota reward
	//
	// - **addSlsCapacity**: log analysis storage capacity reward.
	//
	// example:
	//
	// addAntiRansomwareCapacity
	Reward *string `json:"Reward,omitempty" xml:"Reward,omitempty"`
	// The reward configuration information. This parameter is in JSON format.
	//
	// > The key in the JSON object indicates the reward content, and the value indicates the reward amount. Valid values of the key:
	//
	// - **webLockAuthCount**: the web tamper-proofing authorization quota
	//
	// - **ransomwareCapacity**: the anti-ransomware capacity, in GB
	//
	// - **slsCapacity**: the log analysis capacity, in GB
	//
	// - **days**: the number of usage days
	//
	// - **imageScanAuthCount**: the image scan authorization quota
	//
	// - **honeypotAuthCount**: the cloud honeypot authorization quota.
	//
	// example:
	//
	// {"days":60,"ransomwareCapacity":10}
	RewardConfig *string `json:"RewardConfig,omitempty" xml:"RewardConfig,omitempty"`
}

func (s QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) String() string {
	return dara.Prettify(s)
}

func (s QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) GoString() string {
	return s.String()
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) GetIsRewardTaked() *string {
	return s.IsRewardTaked
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) GetReward() *string {
	return s.Reward
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) GetRewardConfig() *string {
	return s.RewardConfig
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) SetIsRewardTaked(v string) *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData {
	s.IsRewardTaked = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) SetReward(v string) *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData {
	s.Reward = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) SetRewardConfig(v string) *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData {
	s.RewardConfig = &v
	return s
}

func (s *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData) Validate() error {
	return dara.Validate(s)
}
