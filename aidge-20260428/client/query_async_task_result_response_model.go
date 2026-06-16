// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAsyncTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAsyncTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryAsyncTaskResultResponseBody) *QueryAsyncTaskResultResponse
	GetBody() *QueryAsyncTaskResultResponseBody
}

type QueryAsyncTaskResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAsyncTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAsyncTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResultResponse) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAsyncTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAsyncTaskResultResponse) GetBody() *QueryAsyncTaskResultResponseBody {
	return s.Body
}

func (s *QueryAsyncTaskResultResponse) SetHeaders(v map[string]*string) *QueryAsyncTaskResultResponse {
	s.Headers = v
	return s
}

func (s *QueryAsyncTaskResultResponse) SetStatusCode(v int32) *QueryAsyncTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsyncTaskResultResponse) SetBody(v *QueryAsyncTaskResultResponseBody) *QueryAsyncTaskResultResponse {
	s.Body = v
	return s
}

func (s *QueryAsyncTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
