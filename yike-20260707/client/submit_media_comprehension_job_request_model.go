// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaComprehensionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *SubmitMediaComprehensionJobRequest
	GetInput() *string
	SetJobParams(v string) *SubmitMediaComprehensionJobRequest
	GetJobParams() *string
	SetJobType(v string) *SubmitMediaComprehensionJobRequest
	GetJobType() *string
	SetUserData(v string) *SubmitMediaComprehensionJobRequest
	GetUserData() *string
}

type SubmitMediaComprehensionJobRequest struct {
	// The input material. JSON string with the following structure:
	//
	// - Medias (Array<Object>, required): The list of media assets. Contains 1 to 10 elements. Each element includes the following fields:
	//
	//   - Type (String, required): The media asset type. Valid values: video or image (case-insensitive).
	//
	//   - Url (String, either Url or MediaId is required): The URL of the media asset. The URL must start with http:// or https:// and cannot exceed 2048 characters in length. Unregistered URLs are automatically registered as media assets.
	//
	//   - MediaId (String, either Url or MediaId is required): The ID of a registered media asset. If both Url and MediaId are specified, MediaId takes precedence.
	//
	// example:
	//
	// {"Medias":[{"Type":"video","Url":"https://xxx.mp4"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The analysis parameters. JSON string. The total length cannot exceed 65536 characters, and the total number of fields cannot exceed 20.
	//
	// example:
	//
	// {"ProductName":"Quiet Blender Soymilk Maker","BrandName":"LiangChu","SellingPoints":["Low-noise blending","One-touch self-cleaning"]}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The job type.
	//
	// - VideoBreakdown: viral video breakdown. Requires Medias to contain exactly 1 element with Type=video.
	//
	// - ProductRecognition: product image information recognition. Requires all elements in Medias to have Type=image.
	//
	// example:
	//
	// VideoBreakdown
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The custom parameters. JSON string that is returned as-is in the callback result (for example, newsKey). The system reserved field NotifyAddress specifies the callback URL. The callback is triggered after the job is completed.
	//
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaComprehensionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaComprehensionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaComprehensionJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitMediaComprehensionJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitMediaComprehensionJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitMediaComprehensionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaComprehensionJobRequest) SetInput(v string) *SubmitMediaComprehensionJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetJobParams(v string) *SubmitMediaComprehensionJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetJobType(v string) *SubmitMediaComprehensionJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetUserData(v string) *SubmitMediaComprehensionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) Validate() error {
	return dara.Validate(s)
}
