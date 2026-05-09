// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskConcurrencyResponseBody) *QueryTaskConcurrencyResponse
	GetBody() *QueryTaskConcurrencyResponseBody
}

type QueryTaskConcurrencyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskConcurrencyResponse) GetBody() *QueryTaskConcurrencyResponseBody {
	return s.Body
}

func (s *QueryTaskConcurrencyResponse) SetHeaders(v map[string]*string) *QueryTaskConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskConcurrencyResponse) SetStatusCode(v int32) *QueryTaskConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskConcurrencyResponse) SetBody(v *QueryTaskConcurrencyResponseBody) *QueryTaskConcurrencyResponse {
	s.Body = v
	return s
}

func (s *QueryTaskConcurrencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
