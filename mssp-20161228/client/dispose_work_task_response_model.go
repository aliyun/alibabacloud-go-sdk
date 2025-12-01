// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeWorkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisposeWorkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisposeWorkTaskResponse
	GetStatusCode() *int32
	SetBody(v *DisposeWorkTaskResponseBody) *DisposeWorkTaskResponse
	GetBody() *DisposeWorkTaskResponseBody
}

type DisposeWorkTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeWorkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeWorkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DisposeWorkTaskResponse) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisposeWorkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisposeWorkTaskResponse) GetBody() *DisposeWorkTaskResponseBody {
	return s.Body
}

func (s *DisposeWorkTaskResponse) SetHeaders(v map[string]*string) *DisposeWorkTaskResponse {
	s.Headers = v
	return s
}

func (s *DisposeWorkTaskResponse) SetStatusCode(v int32) *DisposeWorkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponse) SetBody(v *DisposeWorkTaskResponseBody) *DisposeWorkTaskResponse {
	s.Body = v
	return s
}

func (s *DisposeWorkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
