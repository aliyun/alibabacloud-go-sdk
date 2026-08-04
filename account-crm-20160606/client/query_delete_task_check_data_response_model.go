// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeleteTaskCheckDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeleteTaskCheckDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeleteTaskCheckDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeleteTaskCheckDataResponseBody) *QueryDeleteTaskCheckDataResponse
	GetBody() *QueryDeleteTaskCheckDataResponseBody
}

type QueryDeleteTaskCheckDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeleteTaskCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeleteTaskCheckDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeleteTaskCheckDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDeleteTaskCheckDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeleteTaskCheckDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeleteTaskCheckDataResponse) GetBody() *QueryDeleteTaskCheckDataResponseBody {
	return s.Body
}

func (s *QueryDeleteTaskCheckDataResponse) SetHeaders(v map[string]*string) *QueryDeleteTaskCheckDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDeleteTaskCheckDataResponse) SetStatusCode(v int32) *QueryDeleteTaskCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeleteTaskCheckDataResponse) SetBody(v *QueryDeleteTaskCheckDataResponseBody) *QueryDeleteTaskCheckDataResponse {
	s.Body = v
	return s
}

func (s *QueryDeleteTaskCheckDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
