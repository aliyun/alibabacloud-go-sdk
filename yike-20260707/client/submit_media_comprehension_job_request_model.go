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
	// The input material. It is a JSON string with the following structure:
	//
	// - Medias (Array<Object>, required): The list of media assets. It contains 1 to 10 elements. Each element includes the following fields:
	//
	//   - Type (String, required): The type of the media asset. Valid values: video and image. The value is not case-sensitive.
	//
	//   - Url (String, mutually exclusive with MediaId): The URL of the media asset. It must start with http:// or https:// and be up to 2,048 characters in length. Unregistered URLs are automatically registered as media assets.
	//
	//   - MediaId (String, mutually exclusive with Url): The ID of the registered media asset. If both Url and MediaId are specified, MediaId takes precedence.
	//
	// example:
	//
	// {"Medias":[{"Type":"video","Url":"https://xxx.mp4"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The analysis parameters. It is a JSON string with a maximum length of 65,536 characters and a maximum of 20 fields.
	//
	// example:
	//
	// {"ProductName":"Quiet Blender Soymilk Maker","BrandName":"LiangChu","SellingPoints":["Low-noise blending","One-touch self-cleaning"]}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The job type. Valid values:
	//
	// - VideoBreakdown: viral video breakdown. The Medias array must contain exactly one element with Type set to video.
	//
	// - ProductRecognition: product image information recognition. All elements in the Medias array must have Type set to image.
	//
	// example:
	//
	// VideoBreakdown
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The custom parameters. It is a JSON string that is returned as-is in the callback result, such as newsKey. NotifyAddress is a system-reserved field for the callback URL. A callback is triggered to this URL after the job is completed.
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
