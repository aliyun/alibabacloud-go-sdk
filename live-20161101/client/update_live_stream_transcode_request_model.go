// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateLiveStreamTranscodeRequest
	GetApp() *string
	SetDomain(v string) *UpdateLiveStreamTranscodeRequest
	GetDomain() *string
	SetEncryptParameters(v string) *UpdateLiveStreamTranscodeRequest
	GetEncryptParameters() *string
	SetLazy(v string) *UpdateLiveStreamTranscodeRequest
	GetLazy() *string
	SetOwnerId(v int64) *UpdateLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveStreamTranscodeRequest
	GetRegionId() *string
	SetTemplate(v string) *UpdateLiveStreamTranscodeRequest
	GetTemplate() *string
}

type UpdateLiveStreamTranscodeRequest struct {
	// The name of the application to which the stream belongs, and it cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// Streamer domain name, not modifiable.
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
	// Specifies whether to enable triggered transcoding. Valid values:
	//
	// 	- **yes**: enables triggered transcoding.
	//
	// 	- **no**: disables triggered transcoding.
	//
	// example:
	//
	// yes
	Lazy     *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Transcoding template, not modifiable.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsd
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s UpdateLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *UpdateLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateLiveStreamTranscodeRequest) GetEncryptParameters() *string {
	return s.EncryptParameters
}

func (s *UpdateLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *UpdateLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateLiveStreamTranscodeRequest) SetApp(v string) *UpdateLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetDomain(v string) *UpdateLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetEncryptParameters(v string) *UpdateLiveStreamTranscodeRequest {
	s.EncryptParameters = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetLazy(v string) *UpdateLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetOwnerId(v int64) *UpdateLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetRegionId(v string) *UpdateLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) SetTemplate(v string) *UpdateLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *UpdateLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
