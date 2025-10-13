// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHistoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHistoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHistoryJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHistoryJobResponseBody) *DeleteHistoryJobResponse
	GetBody() *DeleteHistoryJobResponseBody
}

type DeleteHistoryJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHistoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteHistoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHistoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHistoryJobResponse) GetBody() *DeleteHistoryJobResponseBody {
	return s.Body
}

func (s *DeleteHistoryJobResponse) SetHeaders(v map[string]*string) *DeleteHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteHistoryJobResponse) SetStatusCode(v int32) *DeleteHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHistoryJobResponse) SetBody(v *DeleteHistoryJobResponseBody) *DeleteHistoryJobResponse {
	s.Body = v
	return s
}

func (s *DeleteHistoryJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
