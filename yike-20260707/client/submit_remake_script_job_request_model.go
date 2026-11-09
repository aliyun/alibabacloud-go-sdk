// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRemakeScriptJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemakeParams(v string) *SubmitRemakeScriptJobRequest
	GetRemakeParams() *string
	SetRemakeType(v string) *SubmitRemakeScriptJobRequest
	GetRemakeType() *string
	SetUserData(v string) *SubmitRemakeScriptJobRequest
	GetUserData() *string
}

type SubmitRemakeScriptJobRequest struct {
	// The remake parameters. This parameter is a JSON string. The structure varies based on RemakeType. For faithful-remake, the structure is as follows:
	//
	// - ComprehensionResult (string, required): The content understanding result. This is the URL of the JSON file.
	//
	// - Product (object, required): The product information, including the original product and the new product.
	//
	//   - OriginalProductName (string, required): The product or brand name in the original video to be replaced. This is used to locate and replace mentions of the original product in the script.
	//
	//   - NewProduct (object, required): The new product information.
	//
	//     - ProductName (string, required): The product or brand name.
	//
	//     - Description (string, required): The product description.
	//
	//     - ProductKnowledge (string, required): The physical knowledge of the product, such as the material and usage instructions.
	//
	//     - ProductImages (array of strings, required): The images of the new product. This array cannot be empty. Each item must be an HTTP or HTTPS URL.
	//
	//     - SellingPoints (array of strings, optional): The list of product selling points.
	//
	//     - OriginalPrice (string, optional): The original price.
	//
	//     - CurrentPrice (string, optional): The current price.
	//
	//     - Discount (string, optional): The discount information, such as "50 off for purchases over 200" or "buy one get one free".
	//
	// - Avatar (object, required): The avatar information, including the original avatar and the new avatar.
	//
	//     - NewAvatarImages (array of strings, required): The list of images for the new avatar. This array cannot be empty. Each item must be a media asset ID or an image URL. If a URL is provided, the system automatically registers it as a media asset. **Currently, only one image is supported*	- (only the first element of the array is used). The array format is reserved for future extensions.
	//
	//     - OriginalAvatarName (string, required): The avatar name in the original video to be replaced. This is used to locate and replace the original avatar in the script or video frames.
	//
	// - VoiceoverLanguage (string, optional): The voiceover language. Valid values: `zh` (Chinese), `en` (English), `es` (Spanish), `pt` (Portuguese), `fr` (French), `de` (German), `ja` (Japanese), `ko` (Korean), and `ar` (Arabic). Default value: `zh`.
	//
	// example:
	//
	// {
	//
	//     "ComprehensionResult": "http://xxxx.json",
	//
	//     "Product":
	//
	//     {
	//
	//         "OriginalProductName": "xxxx",
	//
	//         "NewProduct":
	//
	//         {
	//
	//             "ProductName": "xxxx",
	//
	//             "Description": "xxxx",
	//
	//             "ProductImages":
	//
	//             [
	//
	//                 "https://xxxx.png",
	//
	//                 "https://xxxx.png",
	//
	//                 "https://xxxx.png"
	//
	//             ],
	//
	//             "ProductKnowledge": "xxxx"
	//
	//         }
	//
	//     },
	//
	//     "Avatar":
	//
	//     {
	//
	//         "NewAvatarImages":
	//
	//         [
	//
	//             "https://xxxx.png"
	//
	//         ],
	//
	//         "OriginalAvatarName": "xxxx"
	//
	//     },
	//
	//     "VoiceoverLanguage": "zh"
	//
	// }
	RemakeParams *string `json:"RemakeParams,omitempty" xml:"RemakeParams,omitempty"`
	// The remake type. This parameter determines the structure of RemakeParams. Currently, only faithful-remake is supported. faithful-remake indicates a faithful remake that replicates the original viral structure paragraph by paragraph and replaces the product and avatar.
	//
	// example:
	//
	// faithful-remake
	RemakeType *string `json:"RemakeType,omitempty" xml:"RemakeType,omitempty"`
	// The custom parameters. This parameter is a JSON string and is returned as-is in the callback result, such as newsKey. NotifyAddress is a system-reserved field for the callback URL. The system sends a callback to this URL after the task is completed.
	//
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitRemakeScriptJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitRemakeScriptJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitRemakeScriptJobRequest) GetRemakeParams() *string {
	return s.RemakeParams
}

func (s *SubmitRemakeScriptJobRequest) GetRemakeType() *string {
	return s.RemakeType
}

func (s *SubmitRemakeScriptJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitRemakeScriptJobRequest) SetRemakeParams(v string) *SubmitRemakeScriptJobRequest {
	s.RemakeParams = &v
	return s
}

func (s *SubmitRemakeScriptJobRequest) SetRemakeType(v string) *SubmitRemakeScriptJobRequest {
	s.RemakeType = &v
	return s
}

func (s *SubmitRemakeScriptJobRequest) SetUserData(v string) *SubmitRemakeScriptJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitRemakeScriptJobRequest) Validate() error {
	return dara.Validate(s)
}
