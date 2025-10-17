// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDtsJobResponseBody) *DeleteDtsJobResponse
	GetBody() *DeleteDtsJobResponseBody
}

type DeleteDtsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDtsJobResponse) GetBody() *DeleteDtsJobResponseBody {
	return s.Body
}

func (s *DeleteDtsJobResponse) SetHeaders(v map[string]*string) *DeleteDtsJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDtsJobResponse) SetStatusCode(v int32) *DeleteDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDtsJobResponse) SetBody(v *DeleteDtsJobResponseBody) *DeleteDtsJobResponse {
	s.Body = v
	return s
}

func (s *DeleteDtsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
