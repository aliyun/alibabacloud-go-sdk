// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *AsyncTask) *GetAsyncTaskResponse
	GetBody() *AsyncTask
}

type GetAsyncTaskResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncTask         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncTaskResponse) GetBody() *AsyncTask {
	return s.Body
}

func (s *GetAsyncTaskResponse) SetHeaders(v map[string]*string) *GetAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskResponse) SetStatusCode(v int32) *GetAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTaskResponse) SetBody(v *AsyncTask) *GetAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
