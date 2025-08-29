// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFeatureConsistencyCheckJobReplayLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncFeatureConsistencyCheckJobReplayLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncFeatureConsistencyCheckJobReplayLogResponse
	GetStatusCode() *int32
	SetBody(v *SyncFeatureConsistencyCheckJobReplayLogResponseBody) *SyncFeatureConsistencyCheckJobReplayLogResponse
	GetBody() *SyncFeatureConsistencyCheckJobReplayLogResponseBody
}

type SyncFeatureConsistencyCheckJobReplayLogResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncFeatureConsistencyCheckJobReplayLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponse) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) GetBody() *SyncFeatureConsistencyCheckJobReplayLogResponseBody {
	return s.Body
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetHeaders(v map[string]*string) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.Headers = v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetStatusCode(v int32) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetBody(v *SyncFeatureConsistencyCheckJobReplayLogResponseBody) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.Body = v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) Validate() error {
	return dara.Validate(s)
}
