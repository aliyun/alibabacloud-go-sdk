// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditTTSVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditionUrl(v string) *AuditTTSVoiceResponseBody
	GetAuditionUrl() *string
	SetRequestId(v string) *AuditTTSVoiceResponseBody
	GetRequestId() *string
}

type AuditTTSVoiceResponseBody struct {
	// example:
	//
	// http://voicenavigator-cn-shanghai.oss-cn-shanghai.aliyuncs.com/exported_files/2020-02-20/ttsConfig-1582188148528-abd8e407de0a49b381bb591bd91fc073.wav?Expires=1582188208&OSSAccessKeyId=LTAIppQY5rofntVJ&Signature=FaBassElzqGEB0H2TvTKPJsOJHs%3D
	AuditionUrl *string `json:"AuditionUrl,omitempty" xml:"AuditionUrl,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuditTTSVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuditTTSVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceResponseBody) GetAuditionUrl() *string {
	return s.AuditionUrl
}

func (s *AuditTTSVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuditTTSVoiceResponseBody) SetAuditionUrl(v string) *AuditTTSVoiceResponseBody {
	s.AuditionUrl = &v
	return s
}

func (s *AuditTTSVoiceResponseBody) SetRequestId(v string) *AuditTTSVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditTTSVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
