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
	// The list of beginner tasks.
	GuideTaskConfigList []*QueryGuidTaskListResponseBodyGuideTaskConfigList `json:"GuideTaskConfigList,omitempty" xml:"GuideTaskConfigList,omitempty" type:"Repeated"`
	// The request ID.
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
	// The information about the reward for a complete task.
	RewardData *QueryGuidTaskListResponseBodyGuideTaskConfigListRewardData `json:"RewardData,omitempty" xml:"RewardData,omitempty" type:"Struct"`
	// The security score that is increased after you complete the task.
	//
	// example:
	//
	// 80
	SecurityScore *int32 `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
	// The status of the beginner task. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: in progress.
	//
	// 	- **2**: complete.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the beginner task.
	//
	// example:
	//
	// t-000d8slfgx4p40kb64ad
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task type. Valid values:
	//
	// 	- **guid_task_security_score_promote_video**: the task of viewing a video tutorial for beginners.
	//
	// 	- **guide_sub_task_config_defence_hbr**: the task of configuring anti-ransomware for servers.
	//
	// 	- **guide_sub_task_config_uni_defence_hbr**: the task of configuring anti-ransomware for databases.
	//
	// 	- **guid_task_log_analysis_config**: the task of configuring log analysis.
	//
	// 	- **guide_sub_task_web_lock_config**: the task of configuring web tamper proofing.
	//
	// 	- **guide_sub_task_config_anti_crack**: the task of configuring protection against brute-force attacks.
	//
	// 	- **guid_task_container_security_video**: the task of viewing the video on how to protect containers.
	//
	// 	- **guid_task_container_image_scan_config**: the task of configuring container image scan.
	//
	// 	- **guid_task_k8s_log_analysis_config**: the task of configuring threat detection on Kubernetes containers.
	//
	// 	- **guid_task_container_network**: the task of configuring container network visualization.
	//
	// 	- **guide_sub_task_config_add_collection**: the task of saving a console URL.
	//
	// 	- **guide_sub_task_vul_scan**: the task of scanning for vulnerabilities.
	//
	// 	- **guide_sub_task_virusKill**: the task of configuring virus detection and removal.
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
	// Indicates whether the reward is claimed. Valid values:
	//
	// 	- **1**: no.
	//
	// 	- **2**: yes.
	//
	// example:
	//
	// 1
	IsRewardTaked *string `json:"IsRewardTaked,omitempty" xml:"IsRewardTaked,omitempty"`
	// The name of the reward. Valid values:
	//
	// 	- **addTrialDay**: the days of trial use.
	//
	// 	- **addAntiRansomwareCapacity**: the anti-ransomware capacity.
	//
	// 	- **addImageScanAuthCount**: the quota for container image scan.
	//
	// 	- **addWebLockAuthCount**: the quota for web tamper proofing.
	//
	// 	- **addSlsCapacity**: the log storage capacity.
	//
	// example:
	//
	// addWebLockAuthCount
	Reward *string `json:"Reward,omitempty" xml:"Reward,omitempty"`
	// The reward configuration. The value of this parameter is in the JSON format.
	//
	// >  The key indicates the reward type, and the value indicates the number of rewards. Valid values of key:
	//
	// 	- **webLockAuthCount**: the quota for web tamper proofing.
	//
	// 	- **webLockAuthCount**: the anti-ransomware capacity. Unit: GB.
	//
	// 	- **slsCapacity**: the log storage capacity. Unit: GB.
	//
	// 	- **days**: the days of trial use.
	//
	// 	- **imageScanAuthCount**: the quota for container image scan.
	//
	// 	- **honeypotAuthCount**: the quota for cloud honeypot.
	//
	// example:
	//
	// {"days":60,"ransomwareCapacity":100}
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
