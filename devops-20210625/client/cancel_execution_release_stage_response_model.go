// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionReleaseStageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelExecutionReleaseStageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelExecutionReleaseStageResponse
	GetStatusCode() *int32
	SetBody(v *CancelExecutionReleaseStageResponseBody) *CancelExecutionReleaseStageResponse
	GetBody() *CancelExecutionReleaseStageResponseBody
}

type CancelExecutionReleaseStageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelExecutionReleaseStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelExecutionReleaseStageResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionReleaseStageResponse) GoString() string {
	return s.String()
}

func (s *CancelExecutionReleaseStageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelExecutionReleaseStageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelExecutionReleaseStageResponse) GetBody() *CancelExecutionReleaseStageResponseBody {
	return s.Body
}

func (s *CancelExecutionReleaseStageResponse) SetHeaders(v map[string]*string) *CancelExecutionReleaseStageResponse {
	s.Headers = v
	return s
}

func (s *CancelExecutionReleaseStageResponse) SetStatusCode(v int32) *CancelExecutionReleaseStageResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelExecutionReleaseStageResponse) SetBody(v *CancelExecutionReleaseStageResponseBody) *CancelExecutionReleaseStageResponse {
	s.Body = v
	return s
}

func (s *CancelExecutionReleaseStageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
