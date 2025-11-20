// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHanaDatabaseAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopHanaDatabaseAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopHanaDatabaseAsyncResponse
	GetStatusCode() *int32
	SetBody(v *StopHanaDatabaseAsyncResponseBody) *StopHanaDatabaseAsyncResponse
	GetBody() *StopHanaDatabaseAsyncResponseBody
}

type StopHanaDatabaseAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopHanaDatabaseAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopHanaDatabaseAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s StopHanaDatabaseAsyncResponse) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopHanaDatabaseAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopHanaDatabaseAsyncResponse) GetBody() *StopHanaDatabaseAsyncResponseBody {
	return s.Body
}

func (s *StopHanaDatabaseAsyncResponse) SetHeaders(v map[string]*string) *StopHanaDatabaseAsyncResponse {
	s.Headers = v
	return s
}

func (s *StopHanaDatabaseAsyncResponse) SetStatusCode(v int32) *StopHanaDatabaseAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponse) SetBody(v *StopHanaDatabaseAsyncResponseBody) *StopHanaDatabaseAsyncResponse {
	s.Body = v
	return s
}

func (s *StopHanaDatabaseAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
