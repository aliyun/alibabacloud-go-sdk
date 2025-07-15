// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *AddLiveStreamTranscodeRequest
	GetApp() *string
	SetDomain(v string) *AddLiveStreamTranscodeRequest
	GetDomain() *string
	SetEncryptParameters(v string) *AddLiveStreamTranscodeRequest
	GetEncryptParameters() *string
	SetLazy(v string) *AddLiveStreamTranscodeRequest
	GetLazy() *string
	SetOwnerId(v int64) *AddLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveStreamTranscodeRequest
	GetRegionId() *string
	SetTemplate(v string) *AddLiveStreamTranscodeRequest
	GetTemplate() *string
}

type AddLiveStreamTranscodeRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// 	- The transcoding template takes effect only if the value of this parameter is the same as the application name that is specified in the ingest URL. The name can be up to 256 characters in length and can contain digits, letters, hyphens (-), and underscores (_).
	//
	// 	- You can also set this parameter to an asterisk (\\*). Asterisks (\\*) can match any string, including an empty string.
	//
	// >  If you configure a transcoding template for which App is set to an asterisk (\\*), the transcoding template is used only if no transcoding template for which App is set to the same value as AppName in the ingest URL exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The name of the main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The encryption configuration. The value is a JSON string. The following fields are included in the syntax:
	//
	// 	- EncryptType: the type of the encryption. Set the value to **aliyun**.
	//
	// 	- KmsKeyID: the ID of the CMK in KMS.
	//
	// 	- KmsKeyExpireInterval: the validity period of the CMK. Valid values: **60 to 3600**. Unit: seconds.
	//
	// example:
	//
	// {"EncryptType": "aliyun", "KmsKeyID":"afce5722-81d2-43c3-9930-7601da11****","KmsKeyExpireInterval":"3600"}
	EncryptParameters *string `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty"`
	// Specifies whether to use the load-on-demand mechanism for transcoding. Default value: **yes**.
	//
	// example:
	//
	// yes
	Lazy     *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The transcoding template. Valid values:
	//
	// 	- Standard transcoding template:
	//
	//     	- **lld**: low definition
	//
	//     	- **lsd**: standard definition
	//
	//     	- **lhd**: high definition
	//
	//     	- **lud**: ultra-high definition
	//
	// 	- Narrowband HD™ transcoding template:
	//
	//     	- **ld**: low definition
	//
	//     	- **sd**: standard definition
	//
	//     	- **hd**: high definition
	//
	//     	- **ud**: ultra-high definition
	//
	// This parameter is required.
	//
	// example:
	//
	// lsd
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s AddLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *AddLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddLiveStreamTranscodeRequest) GetEncryptParameters() *string {
	return s.EncryptParameters
}

func (s *AddLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *AddLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *AddLiveStreamTranscodeRequest) SetApp(v string) *AddLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetDomain(v string) *AddLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetEncryptParameters(v string) *AddLiveStreamTranscodeRequest {
	s.EncryptParameters = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetLazy(v string) *AddLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetOwnerId(v int64) *AddLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetRegionId(v string) *AddLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetTemplate(v string) *AddLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
