// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSampleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSampleResponseBody) *DeleteSampleResponse
	GetBody() *DeleteSampleResponseBody
}

type DeleteSampleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSampleResponse) GetBody() *DeleteSampleResponseBody {
	return s.Body
}

func (s *DeleteSampleResponse) SetHeaders(v map[string]*string) *DeleteSampleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSampleResponse) SetStatusCode(v int32) *DeleteSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSampleResponse) SetBody(v *DeleteSampleResponseBody) *DeleteSampleResponse {
	s.Body = v
	return s
}

func (s *DeleteSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
