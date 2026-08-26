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
	// The name of the application to which the stream belongs.
	//
	// - The value of App must be the same as the AppName in the ingest URL for the transcoding template to take effect. The value can be up to 256 characters in length and can contain digits, letters, hyphens (-), and underscores (_).
	//
	// - App also supports a single asterisk (\\*) as the value, which matches any string including an empty string.
	//
	// > If a transcoding template with App set to a single asterisk (\\*) is configured: when a user pulls a transcoded stream, the system first matches the transcoding template whose App value is the same as the AppName in the ingest URL. If no such template exists, the system matches the transcoding template with App set to a single asterisk (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The streaming domain of the streamer.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The encryption configuration. JSON format. The fields are described as follows:
	//
	// - EncryptType: the encryption type. Set the value to **aliyun**.
	//
	// - KmsKeyID: the user KMS master key ID.
	//
	// - KmsKeyExpireInterval: the key rotation interval. Valid values: **60 to 3600**. Unit: seconds.
	//
	// example:
	//
	// {"EncryptType": "aliyun", "KmsKeyID":"afce5722-81d2-43c3-9930-7601da11****","KmsKeyExpireInterval":"3600"}
	EncryptParameters *string `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty"`
	// Specifies whether to enable on-demand transcoding. Valid values:
	//
	// - **yes**: enables on-demand transcoding.
	//
	// - **no**: disables on-demand transcoding.
	//
	// example:
	//
	// yes
	Lazy    *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. When using SDK 2.0, set this parameter to the region ID that corresponds to the service registration endpoint. When using SDK 1.0, ignore this parameter.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The transcoding template. Valid values:
	//
	// - Standard quality templates:
	//
	//     - **lld**: low definition.
	//
	//     - **lsd**: standard definition.
	//
	//     - **lhd**: high definition.
	//
	//     - **lud**: ultra-high definition.
	//
	//
	//
	// - Narrowband HD™ transcoding templates:
	//
	//     - **ld**: low definition.
	//
	//
	//
	//     - **sd**: standard definition.
	//
	//
	//
	//     - **hd**: high definition.
	//
	//
	//
	//     - **ud**: ultra-high definition.
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
