// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneMediaSelectionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfig(v string) *SubmitSceneMediaSelectionJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitSceneMediaSelectionJobRequest
	GetInputConfig() *string
	SetJobType(v string) *SubmitSceneMediaSelectionJobRequest
	GetJobType() *string
	SetOutputConfig(v string) *SubmitSceneMediaSelectionJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitSceneMediaSelectionJobRequest
	GetUserData() *string
}

type SubmitSceneMediaSelectionJobRequest struct {
	// The editing configuration. Its structure depends on the value of JobType.
	//
	// 	- When JobType is set to Smart_Mix_Media_Select, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Media_Select, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
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
	// 	- When JobType is set to Smart_Mix_Media_Select, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Media_Select, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
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
	// 		"Grand opening! A Freshippo store opens today at the nearby mall.",
	//
	// 		"Great deals on snacks and drinks. Stop by!"
	//
	// 	]
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The job type. Valid values:
	//
	// 	- Smart_Mix_Media_Select
	//
	// 	- Screen_Media_Highlights_Media_Select
	//
	// Differences:
	//
	// 	- Smart_Mix_Media_Select: Matches voiceover scripts with provided video/image materials to select the most relevant clips and returns the matching results. Two options are available: Common mode, which is suitable for general-purpose materials like lifestyle vlogs, travel videos, and marketing content; Movie collections, which is optimized for materials with a coherent plot and specific characters, such as TV series and movies.
	//
	// 	- Screen_Media_Highlights_Media_Select: Automatically identifies and selects clips that are exciting or represent key story points from longer video materials.
	//
	// This parameter is required.
	//
	// example:
	//
	// Smart_Mix_Media_Select
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The output configuration. Its structure and required fields depend on the value of JobType.
	//
	// 	- When JobType is set to Smart_Mix_Media_Select, see [Image-text matching](https://help.aliyun.com/zh/ims/use-cases/intelligent-graphic-matching-into-a-piece/?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_1.7c3d6997qndkZj).
	//
	// 	- When JobType is set to Screen_Media_Highlights_Media_Select, see [Highlight mashup](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.84b5661bIcQULE).
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

func (s SubmitSceneMediaSelectionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneMediaSelectionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSceneMediaSelectionJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitSceneMediaSelectionJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitSceneMediaSelectionJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitSceneMediaSelectionJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSceneMediaSelectionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSceneMediaSelectionJobRequest) SetEditingConfig(v string) *SubmitSceneMediaSelectionJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobRequest) SetInputConfig(v string) *SubmitSceneMediaSelectionJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobRequest) SetJobType(v string) *SubmitSceneMediaSelectionJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobRequest) SetOutputConfig(v string) *SubmitSceneMediaSelectionJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobRequest) SetUserData(v string) *SubmitSceneMediaSelectionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobRequest) Validate() error {
	return dara.Validate(s)
}
