// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityWatchTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityWatchTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityWatchTaskLogResponseBody) *GetQualityWatchTaskLogResponse
	GetBody() *GetQualityWatchTaskLogResponseBody
}

type GetQualityWatchTaskLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityWatchTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityWatchTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskLogResponse) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityWatchTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityWatchTaskLogResponse) GetBody() *GetQualityWatchTaskLogResponseBody {
	return s.Body
}

func (s *GetQualityWatchTaskLogResponse) SetHeaders(v map[string]*string) *GetQualityWatchTaskLogResponse {
	s.Headers = v
	return s
}

func (s *GetQualityWatchTaskLogResponse) SetStatusCode(v int32) *GetQualityWatchTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityWatchTaskLogResponse) SetBody(v *GetQualityWatchTaskLogResponseBody) *GetQualityWatchTaskLogResponse {
	s.Body = v
	return s
}

func (s *GetQualityWatchTaskLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
