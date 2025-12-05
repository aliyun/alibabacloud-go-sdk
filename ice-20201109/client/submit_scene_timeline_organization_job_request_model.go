// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneTimelineOrganizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetInputConfig() *string
	SetJobType(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetJobType() *string
	SetMediaSelectResult(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetMediaSelectResult() *string
	SetOutputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetUserData() *string
}

type SubmitSceneTimelineOrganizationJobRequest struct {
	// The editing configuration. Its structure depends on the value of JobType.
	//
	// 	- When JobType is set to Smart_Mix_Timeline_Organize, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Timeline_Organize, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
	//
	// example:
	//
	// {
	//
	//   "MediaConfig": {
	//
	//       "Volume": 0
	//
	//   },
	//
	//   "SpeechConfig": {
	//
	//       "Volume": 1
	//
	//   },
	//
	//  "BackgroundMusicConfig": {
	//
	//       "Volume": 0.3
	//
	//   }
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configuration. Its structure and required fields depend on the value of JobType.
	//
	// 	- When JobType is set to Smart_Mix_Timeline_Organize, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Timeline_Organize, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"BackgroundMusic": "****75c3936f3a8743850f2da942****",
	//
	// 	"MediaArray": [
	//
	// 		"https://test-bucket.oss-cn-shanghai.aliyuncs.com/test.mp4"
	//
	// 	],
	//
	// 	"SpeechTextArray": [
	//
	// 		"A new Freshippo store just opened in the nearby mall. Today is the grand opening."
	//
	// 	]
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The job type. Valid values:
	//
	// 	- Smart_Mix_Timeline_Organize: Image-text matching.
	//
	// 	- Screen_Media_Highlights_Timeline_Organize: Highlight mashup.
	//
	// Differences:
	//
	// 	- Image-text matching: Arranges a timeline based on the results of matching a voiceover script to media assets. Ideal for bulk marketing videos and general-purpose montages.
	//
	// 	- Highlight mashup: Arranges a timeline based on the results of highlight clip selection. Ideal for creating action-packed highlight reels from short-form dramas.
	//
	// This parameter is required.
	//
	// example:
	//
	// Smart_Mix_Timeline_Organize
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The media selection results from a previously run SubmitSceneMediaSelectionJob. You can retrieve this result by calling GetBatchMediaProducingJob.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"textMatchMediaOutputList": [{
	//
	// 		"textMatchMediaSentenceOutputList": [{
	//
	// 				"duration": 3.366667,
	//
	// 				"matchClipList": [{
	//
	// 					"endTime": 11.16,
	//
	// 					"mediaId": "****a0900f5071efbf1ce7e6c66a****",
	//
	// 					"startTime": 8.04
	//
	// 				}],
	//
	// 				"text": "A new Freshippo store just opened in the nearby mall"
	//
	// 			},
	//
	// 			{
	//
	// 				"duration": 1.566667,
	//
	// 				"matchClipList": [{
	//
	// 					"endTime": 1.54,
	//
	// 					"mediaId": "****a0900f5071efbf1ce7e6c66a****",
	//
	// 					"startTime": 0
	//
	// 				}],
	//
	// 				"text": "Today is the grand opening"
	//
	// 			}
	//
	// 		]
	//
	// 	}]
	//
	// }
	MediaSelectResult *string `json:"MediaSelectResult,omitempty" xml:"MediaSelectResult,omitempty"`
	// The output configuration. Its structure and required fields depend on the value of JobType.
	//
	// 	- When JobType is set to Smart_Mix_Timeline_Organize, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Timeline_Organize, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://test-bucket.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 1,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The user-defined data, including the business and callback configurations. For more information, see [UserData](~~357745#section-urj-v3f-0s1~~).
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"} or {"NotifyAddress":"https://xx.xx.xxx"} or {"NotifyAddress":"ice-callback-demo"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSceneTimelineOrganizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneTimelineOrganizationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetMediaSelectResult() *string {
	return s.MediaSelectResult
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetEditingConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetInputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetJobType(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetMediaSelectResult(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.MediaSelectResult = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetOutputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetUserData(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) Validate() error {
	return dara.Validate(s)
}
