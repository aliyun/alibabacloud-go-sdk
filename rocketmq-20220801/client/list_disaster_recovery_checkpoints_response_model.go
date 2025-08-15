// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryCheckpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDisasterRecoveryCheckpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDisasterRecoveryCheckpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListDisasterRecoveryCheckpointsResponseBody) *ListDisasterRecoveryCheckpointsResponse
	GetBody() *ListDisasterRecoveryCheckpointsResponseBody
}

type ListDisasterRecoveryCheckpointsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisasterRecoveryCheckpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponse) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDisasterRecoveryCheckpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDisasterRecoveryCheckpointsResponse) GetBody() *ListDisasterRecoveryCheckpointsResponseBody {
	return s.Body
}

func (s *ListDisasterRecoveryCheckpointsResponse) SetHeaders(v map[string]*string) *ListDisasterRecoveryCheckpointsResponse {
	s.Headers = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponse) SetStatusCode(v int32) *ListDisasterRecoveryCheckpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponse) SetBody(v *ListDisasterRecoveryCheckpointsResponseBody) *ListDisasterRecoveryCheckpointsResponse {
	s.Body = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponse) Validate() error {
	return dara.Validate(s)
}
