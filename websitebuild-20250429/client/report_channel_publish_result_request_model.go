// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportChannelPublishResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelAccountName(v string) *ReportChannelPublishResultRequest
	GetChannelAccountName() *string
	SetDraftId(v string) *ReportChannelPublishResultRequest
	GetDraftId() *string
	SetExternalId(v string) *ReportChannelPublishResultRequest
	GetExternalId() *string
	SetExternalUrl(v string) *ReportChannelPublishResultRequest
	GetExternalUrl() *string
	SetFailReason(v string) *ReportChannelPublishResultRequest
	GetFailReason() *string
	SetSuccess(v bool) *ReportChannelPublishResultRequest
	GetSuccess() *bool
}

type ReportChannelPublishResultRequest struct {
	// The display name of the publishing account. A null value does not overwrite the original value. Maximum length: 256 characters.
	//
	// example:
	//
	// xx
	ChannelAccountName *string `json:"ChannelAccountName,omitempty" xml:"ChannelAccountName,omitempty"`
	// The channel draft ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft-789
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The ID returned by the platform. Set this parameter when the publishing is successful.
	//
	// example:
	//
	// 5467410713875420090
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The redirect URL of the platform.
	//
	// example:
	//
	// https://hf-groot.lingostarcdn.com/sparkenglish-sv/groot-parent/1-698a94b7-133b20ab-42ae793d40f6/05a0caf1-abb0-415d-941c-eecf0b613eae.mp4
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// The failure reason. Set this parameter when the publishing fails.
	//
	// example:
	//
	// xxx
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// Specifies whether the publishing is successful (true/false).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportChannelPublishResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportChannelPublishResultRequest) GoString() string {
	return s.String()
}

func (s *ReportChannelPublishResultRequest) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *ReportChannelPublishResultRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *ReportChannelPublishResultRequest) GetExternalId() *string {
	return s.ExternalId
}

func (s *ReportChannelPublishResultRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *ReportChannelPublishResultRequest) GetFailReason() *string {
	return s.FailReason
}

func (s *ReportChannelPublishResultRequest) GetSuccess() *bool {
	return s.Success
}

func (s *ReportChannelPublishResultRequest) SetChannelAccountName(v string) *ReportChannelPublishResultRequest {
	s.ChannelAccountName = &v
	return s
}

func (s *ReportChannelPublishResultRequest) SetDraftId(v string) *ReportChannelPublishResultRequest {
	s.DraftId = &v
	return s
}

func (s *ReportChannelPublishResultRequest) SetExternalId(v string) *ReportChannelPublishResultRequest {
	s.ExternalId = &v
	return s
}

func (s *ReportChannelPublishResultRequest) SetExternalUrl(v string) *ReportChannelPublishResultRequest {
	s.ExternalUrl = &v
	return s
}

func (s *ReportChannelPublishResultRequest) SetFailReason(v string) *ReportChannelPublishResultRequest {
	s.FailReason = &v
	return s
}

func (s *ReportChannelPublishResultRequest) SetSuccess(v bool) *ReportChannelPublishResultRequest {
	s.Success = &v
	return s
}

func (s *ReportChannelPublishResultRequest) Validate() error {
	return dara.Validate(s)
}
