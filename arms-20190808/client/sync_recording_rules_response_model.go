// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRecordingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncRecordingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncRecordingRulesResponse
	GetStatusCode() *int32
	SetBody(v *SyncRecordingRulesResponseBody) *SyncRecordingRulesResponse
	GetBody() *SyncRecordingRulesResponseBody
}

type SyncRecordingRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncRecordingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncRecordingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncRecordingRulesResponse) GoString() string {
	return s.String()
}

func (s *SyncRecordingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncRecordingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncRecordingRulesResponse) GetBody() *SyncRecordingRulesResponseBody {
	return s.Body
}

func (s *SyncRecordingRulesResponse) SetHeaders(v map[string]*string) *SyncRecordingRulesResponse {
	s.Headers = v
	return s
}

func (s *SyncRecordingRulesResponse) SetStatusCode(v int32) *SyncRecordingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncRecordingRulesResponse) SetBody(v *SyncRecordingRulesResponseBody) *SyncRecordingRulesResponse {
	s.Body = v
	return s
}

func (s *SyncRecordingRulesResponse) Validate() error {
	return dara.Validate(s)
}
