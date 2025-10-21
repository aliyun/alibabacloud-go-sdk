// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentAsyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContentAsyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContentAsyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ContentAsyncDetectResponseBody) *ContentAsyncDetectResponse
	GetBody() *ContentAsyncDetectResponseBody
}

type ContentAsyncDetectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContentAsyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContentAsyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContentAsyncDetectResponse) GetBody() *ContentAsyncDetectResponseBody {
	return s.Body
}

func (s *ContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ContentAsyncDetectResponse) SetStatusCode(v int32) *ContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ContentAsyncDetectResponse) SetBody(v *ContentAsyncDetectResponseBody) *ContentAsyncDetectResponse {
	s.Body = v
	return s
}

func (s *ContentAsyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
