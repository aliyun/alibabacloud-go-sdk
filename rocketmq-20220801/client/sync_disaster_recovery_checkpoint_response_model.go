// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDisasterRecoveryCheckpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDisasterRecoveryCheckpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDisasterRecoveryCheckpointResponse
	GetStatusCode() *int32
	SetBody(v *SyncDisasterRecoveryCheckpointResponseBody) *SyncDisasterRecoveryCheckpointResponse
	GetBody() *SyncDisasterRecoveryCheckpointResponseBody
}

type SyncDisasterRecoveryCheckpointResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDisasterRecoveryCheckpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDisasterRecoveryCheckpointResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDisasterRecoveryCheckpointResponse) GoString() string {
	return s.String()
}

func (s *SyncDisasterRecoveryCheckpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDisasterRecoveryCheckpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDisasterRecoveryCheckpointResponse) GetBody() *SyncDisasterRecoveryCheckpointResponseBody {
	return s.Body
}

func (s *SyncDisasterRecoveryCheckpointResponse) SetHeaders(v map[string]*string) *SyncDisasterRecoveryCheckpointResponse {
	s.Headers = v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponse) SetStatusCode(v int32) *SyncDisasterRecoveryCheckpointResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponse) SetBody(v *SyncDisasterRecoveryCheckpointResponseBody) *SyncDisasterRecoveryCheckpointResponse {
	s.Body = v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponse) Validate() error {
	return dara.Validate(s)
}
