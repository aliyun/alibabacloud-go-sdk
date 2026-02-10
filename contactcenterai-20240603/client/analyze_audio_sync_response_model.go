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
	SetId(v string) *AnalyzeAudioSyncResponse
	GetId() *string
	SetEvent(v string) *AnalyzeAudioSyncResponse
	GetEvent() *string
	SetBody(v *AnalyzeAudioSyncResponseBody) *AnalyzeAudioSyncResponse
	GetBody() *AnalyzeAudioSyncResponseBody
}

type AnalyzeAudioSyncResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *AnalyzeAudioSyncResponse) GetId() *string {
	return s.Id
}

func (s *AnalyzeAudioSyncResponse) GetEvent() *string {
	return s.Event
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

func (s *AnalyzeAudioSyncResponse) SetId(v string) *AnalyzeAudioSyncResponse {
	s.Id = &v
	return s
}

func (s *AnalyzeAudioSyncResponse) SetEvent(v string) *AnalyzeAudioSyncResponse {
	s.Event = &v
	return s
}

func (s *AnalyzeAudioSyncResponse) SetBody(v *AnalyzeAudioSyncResponseBody) *AnalyzeAudioSyncResponse {
	s.Body = v
	return s
}

func (s *AnalyzeAudioSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
