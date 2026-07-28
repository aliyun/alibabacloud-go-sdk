// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutopilotTuningHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutopilotTuningHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutopilotTuningHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutopilotTuningHistoriesResponseBody) *ListAutopilotTuningHistoriesResponse
	GetBody() *ListAutopilotTuningHistoriesResponseBody
}

type ListAutopilotTuningHistoriesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutopilotTuningHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutopilotTuningHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutopilotTuningHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAutopilotTuningHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutopilotTuningHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutopilotTuningHistoriesResponse) GetBody() *ListAutopilotTuningHistoriesResponseBody {
	return s.Body
}

func (s *ListAutopilotTuningHistoriesResponse) SetHeaders(v map[string]*string) *ListAutopilotTuningHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAutopilotTuningHistoriesResponse) SetStatusCode(v int32) *ListAutopilotTuningHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponse) SetBody(v *ListAutopilotTuningHistoriesResponseBody) *ListAutopilotTuningHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListAutopilotTuningHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
