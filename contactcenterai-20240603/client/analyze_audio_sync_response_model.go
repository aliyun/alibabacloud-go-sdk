// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeAudioSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeAudioSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeAudioSyncResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeAudioSyncResponseBody) *AnalyzeAudioSyncResponse
	GetBody() *AnalyzeAudioSyncResponseBody
}

type AnalyzeAudioSyncResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeAudioSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeAudioSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeAudioSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeAudioSyncResponse) GetBody() *AnalyzeAudioSyncResponseBody {
	return s.Body
}

func (s *AnalyzeAudioSyncResponse) SetHeaders(v map[string]*string) *AnalyzeAudioSyncResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeAudioSyncResponse) SetStatusCode(v int32) *AnalyzeAudioSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeAudioSyncResponse) SetBody(v *AnalyzeAudioSyncResponseBody) *AnalyzeAudioSyncResponse {
	s.Body = v
	return s
}

func (s *AnalyzeAudioSyncResponse) Validate() error {
	return dara.Validate(s)
}
