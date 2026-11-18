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
	// The input media. The value is a JSON string with the following structure:
	//
	// - Medias (Array<Object>, required): The list of media files. The array must contain 1 to 10 elements. Each element contains the following fields:
	//
	//   - Type (String, required): The media type. Valid values: video and image. The value is not case-sensitive.
	//
	//   - Url (String, mutually exclusive with MediaId): The URL of the media file. The URL must start with http:// or https:// and be up to 2,048 characters in length. Unregistered URLs are automatically registered as media files.
	//
	//   - MediaId (String, mutually exclusive with Url): The ID of the registered media file. If both Url and MediaId are specified, MediaId takes precedence.
	//
	// example:
	//
	// {
	//
	//   "Medias": [
	//
	//     {
	//
	//       "Type": "video",
	//
	//       "Url": "https://example.com/source.mp4"
	//
	//     }
	//
	//   ]
	//
	// }
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The analysis parameters. The value is a JSON string up to 65,536 characters in length.
	//
	// - VideoBreakdown: Supports the ResultType parameter (String, optional). Valid value: Standard (default).
	//
	// - ProductRecognition: Uses parameters such as product information and prompts. The total number of fields cannot exceed 20. Example fields for product parameters include ProductName, BrandName, and SellingPoints.
	//
	// example:
	//
	// {
	//
	//   "ResultType": "Standard"
	//
	// }
	//
	// or
	//
	// {
	//
	//   "ProductName": "Quiet Blender Soymilk Maker",
	//
	//   "BrandName": "LiangChu",
	//
	//   "SellingPoints": [
	//
	//     "Low-noise blending",
	//
	//     "One-touch self-cleaning"
	//
	//   ]
	//
	// }
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The job type. Valid values:
	//
	// - VideoBreakdown: Viral video breakdown. The Medias array must contain exactly one element with Type set to video.
	//
	// - ProductRecognition: Product image information recognition. All elements in the Medias array must have Type set to image.
	//
	// example:
	//
	// VideoBreakdown
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The custom parameters. The value is a JSON string. The custom parameters are returned as-is in the callback result (for example, newsKey). NotifyAddress is a system-reserved field for the callback URL. A callback is triggered after the job is completed.
	//
	// example:
	//
	// {
	//
	//   "NotifyAddress": "https://example.com/callback"
	//
	// }
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
