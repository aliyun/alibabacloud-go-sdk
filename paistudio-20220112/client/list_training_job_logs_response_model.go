// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobLogsResponseBody) *ListTrainingJobLogsResponse
	GetBody() *ListTrainingJobLogsResponseBody
}

type ListTrainingJobLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobLogsResponse) GetBody() *ListTrainingJobLogsResponseBody {
	return s.Body
}

func (s *ListTrainingJobLogsResponse) SetHeaders(v map[string]*string) *ListTrainingJobLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobLogsResponse) SetStatusCode(v int32) *ListTrainingJobLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobLogsResponse) SetBody(v *ListTrainingJobLogsResponseBody) *ListTrainingJobLogsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobLogsResponse) Validate() error {
	return dara.Validate(s)
}
