// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStreamTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStreamTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetStreamTaskResponseBody) *GetStreamTaskResponse
	GetBody() *GetStreamTaskResponseBody
}

type GetStreamTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStreamTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStreamTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTaskResponse) GoString() string {
	return s.String()
}

func (s *GetStreamTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStreamTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStreamTaskResponse) GetBody() *GetStreamTaskResponseBody {
	return s.Body
}

func (s *GetStreamTaskResponse) SetHeaders(v map[string]*string) *GetStreamTaskResponse {
	s.Headers = v
	return s
}

func (s *GetStreamTaskResponse) SetStatusCode(v int32) *GetStreamTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStreamTaskResponse) SetBody(v *GetStreamTaskResponseBody) *GetStreamTaskResponse {
	s.Body = v
	return s
}

func (s *GetStreamTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
