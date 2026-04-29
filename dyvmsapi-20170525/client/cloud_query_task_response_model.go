// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryTaskResponseBody) *CloudQueryTaskResponse
	GetBody() *CloudQueryTaskResponseBody
}

type CloudQueryTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryTaskResponse) GetBody() *CloudQueryTaskResponseBody {
	return s.Body
}

func (s *CloudQueryTaskResponse) SetHeaders(v map[string]*string) *CloudQueryTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryTaskResponse) SetStatusCode(v int32) *CloudQueryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryTaskResponse) SetBody(v *CloudQueryTaskResponseBody) *CloudQueryTaskResponse {
	s.Body = v
	return s
}

func (s *CloudQueryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
