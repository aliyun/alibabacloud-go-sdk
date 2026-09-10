// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerReadAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInnerReadAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInnerReadAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetInnerReadAsyncResultResponseBody) *GetInnerReadAsyncResultResponse
	GetBody() *GetInnerReadAsyncResultResponseBody
}

type GetInnerReadAsyncResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInnerReadAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInnerReadAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInnerReadAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetInnerReadAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInnerReadAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInnerReadAsyncResultResponse) GetBody() *GetInnerReadAsyncResultResponseBody {
	return s.Body
}

func (s *GetInnerReadAsyncResultResponse) SetHeaders(v map[string]*string) *GetInnerReadAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetInnerReadAsyncResultResponse) SetStatusCode(v int32) *GetInnerReadAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInnerReadAsyncResultResponse) SetBody(v *GetInnerReadAsyncResultResponseBody) *GetInnerReadAsyncResultResponse {
	s.Body = v
	return s
}

func (s *GetInnerReadAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
