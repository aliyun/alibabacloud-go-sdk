// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityWatchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitQualityWatchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitQualityWatchTasksResponse
	GetStatusCode() *int32
	SetBody(v *SubmitQualityWatchTasksResponseBody) *SubmitQualityWatchTasksResponse
	GetBody() *SubmitQualityWatchTasksResponseBody
}

type SubmitQualityWatchTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitQualityWatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitQualityWatchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitQualityWatchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitQualityWatchTasksResponse) GetBody() *SubmitQualityWatchTasksResponseBody {
	return s.Body
}

func (s *SubmitQualityWatchTasksResponse) SetHeaders(v map[string]*string) *SubmitQualityWatchTasksResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualityWatchTasksResponse) SetStatusCode(v int32) *SubmitQualityWatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualityWatchTasksResponse) SetBody(v *SubmitQualityWatchTasksResponseBody) *SubmitQualityWatchTasksResponse {
	s.Body = v
	return s
}

func (s *SubmitQualityWatchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
