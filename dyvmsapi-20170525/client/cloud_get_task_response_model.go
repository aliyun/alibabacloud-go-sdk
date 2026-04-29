// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetTaskResponseBody) *CloudGetTaskResponse
	GetBody() *CloudGetTaskResponseBody
}

type CloudGetTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudGetTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetTaskResponse) GetBody() *CloudGetTaskResponseBody {
	return s.Body
}

func (s *CloudGetTaskResponse) SetHeaders(v map[string]*string) *CloudGetTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudGetTaskResponse) SetStatusCode(v int32) *CloudGetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetTaskResponse) SetBody(v *CloudGetTaskResponseBody) *CloudGetTaskResponse {
	s.Body = v
	return s
}

func (s *CloudGetTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
