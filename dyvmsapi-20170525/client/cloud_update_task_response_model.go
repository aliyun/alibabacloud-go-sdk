// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudUpdateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudUpdateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudUpdateTaskResponseBody) *CloudUpdateTaskResponse
	GetBody() *CloudUpdateTaskResponseBody
}

type CloudUpdateTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudUpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudUpdateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudUpdateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudUpdateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudUpdateTaskResponse) GetBody() *CloudUpdateTaskResponseBody {
	return s.Body
}

func (s *CloudUpdateTaskResponse) SetHeaders(v map[string]*string) *CloudUpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudUpdateTaskResponse) SetStatusCode(v int32) *CloudUpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudUpdateTaskResponse) SetBody(v *CloudUpdateTaskResponseBody) *CloudUpdateTaskResponse {
	s.Body = v
	return s
}

func (s *CloudUpdateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
