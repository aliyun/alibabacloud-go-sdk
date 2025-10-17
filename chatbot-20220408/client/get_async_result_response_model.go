// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncResultResponseBody) *GetAsyncResultResponse
	GetBody() *GetAsyncResultResponseBody
}

type GetAsyncResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncResultResponse) GetBody() *GetAsyncResultResponseBody {
	return s.Body
}

func (s *GetAsyncResultResponse) SetHeaders(v map[string]*string) *GetAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncResultResponse) SetStatusCode(v int32) *GetAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncResultResponse) SetBody(v *GetAsyncResultResponseBody) *GetAsyncResultResponse {
	s.Body = v
	return s
}

func (s *GetAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
