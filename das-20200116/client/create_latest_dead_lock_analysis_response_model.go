// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLatestDeadLockAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLatestDeadLockAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLatestDeadLockAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *CreateLatestDeadLockAnalysisResponseBody) *CreateLatestDeadLockAnalysisResponse
	GetBody() *CreateLatestDeadLockAnalysisResponseBody
}

type CreateLatestDeadLockAnalysisResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLatestDeadLockAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLatestDeadLockAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLatestDeadLockAnalysisResponse) GoString() string {
	return s.String()
}

func (s *CreateLatestDeadLockAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLatestDeadLockAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLatestDeadLockAnalysisResponse) GetBody() *CreateLatestDeadLockAnalysisResponseBody {
	return s.Body
}

func (s *CreateLatestDeadLockAnalysisResponse) SetHeaders(v map[string]*string) *CreateLatestDeadLockAnalysisResponse {
	s.Headers = v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponse) SetStatusCode(v int32) *CreateLatestDeadLockAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponse) SetBody(v *CreateLatestDeadLockAnalysisResponseBody) *CreateLatestDeadLockAnalysisResponse {
	s.Body = v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
