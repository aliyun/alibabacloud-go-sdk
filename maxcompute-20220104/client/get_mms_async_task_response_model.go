// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsAsyncTaskResponseBody) *GetMmsAsyncTaskResponse
	GetBody() *GetMmsAsyncTaskResponseBody
}

type GetMmsAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsAsyncTaskResponse) GetBody() *GetMmsAsyncTaskResponseBody {
	return s.Body
}

func (s *GetMmsAsyncTaskResponse) SetHeaders(v map[string]*string) *GetMmsAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMmsAsyncTaskResponse) SetStatusCode(v int32) *GetMmsAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsAsyncTaskResponse) SetBody(v *GetMmsAsyncTaskResponseBody) *GetMmsAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetMmsAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
