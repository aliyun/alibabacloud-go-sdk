// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PredImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PredImageResponse
	GetStatusCode() *int32
	SetBody(v *PredImageResponseBody) *PredImageResponse
	GetBody() *PredImageResponseBody
}

type PredImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredImageResponse) String() string {
	return dara.Prettify(s)
}

func (s PredImageResponse) GoString() string {
	return s.String()
}

func (s *PredImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PredImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PredImageResponse) GetBody() *PredImageResponseBody {
	return s.Body
}

func (s *PredImageResponse) SetHeaders(v map[string]*string) *PredImageResponse {
	s.Headers = v
	return s
}

func (s *PredImageResponse) SetStatusCode(v int32) *PredImageResponse {
	s.StatusCode = &v
	return s
}

func (s *PredImageResponse) SetBody(v *PredImageResponseBody) *PredImageResponse {
	s.Body = v
	return s
}

func (s *PredImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
