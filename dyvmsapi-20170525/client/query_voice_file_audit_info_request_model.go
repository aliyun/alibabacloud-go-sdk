// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVoiceFileAuditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *QueryVoiceFileAuditInfoRequest
	GetBusinessType() *int32
	SetOwnerId(v int64) *QueryVoiceFileAuditInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryVoiceFileAuditInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVoiceFileAuditInfoRequest
	GetResourceOwnerId() *int64
	SetVoiceCodes(v string) *QueryVoiceFileAuditInfoRequest
	GetVoiceCodes() *string
}

type QueryVoiceFileAuditInfoRequest struct {
	// The type of the voice file. Valid values:
	//
	// 	- **0*	- (default): the voice notification file.
	//
	// 	- **2**: the recording file.
	//
	// example:
	//
	// 0
	BusinessType         *int32  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the voice file. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages*	- > **Voice Notifications*	- or **Voice File Management**, and then click the **Voice Notification Files*	- tab to view the **voice ID**.
	//
	// > You can query up to 10 voice files each time. Separate the voice file names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 8501d2eb-efbb-471f-xxx8-****.wav
	VoiceCodes *string `json:"VoiceCodes,omitempty" xml:"VoiceCodes,omitempty"`
}

func (s QueryVoiceFileAuditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVoiceFileAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *QueryVoiceFileAuditInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVoiceFileAuditInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVoiceFileAuditInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVoiceFileAuditInfoRequest) GetVoiceCodes() *string {
	return s.VoiceCodes
}

func (s *QueryVoiceFileAuditInfoRequest) SetBusinessType(v int32) *QueryVoiceFileAuditInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetOwnerId(v int64) *QueryVoiceFileAuditInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetResourceOwnerAccount(v string) *QueryVoiceFileAuditInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetResourceOwnerId(v int64) *QueryVoiceFileAuditInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetVoiceCodes(v string) *QueryVoiceFileAuditInfoRequest {
	s.VoiceCodes = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) Validate() error {
	return dara.Validate(s)
}
