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
	// example:
	//
	// faithful-remake
	RemakeType *string `json:"RemakeType,omitempty" xml:"RemakeType,omitempty"`
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
