// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteTaskResponseBody) *CloudDeleteTaskResponse
	GetBody() *CloudDeleteTaskResponseBody
}

type CloudDeleteTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteTaskResponse) GetBody() *CloudDeleteTaskResponseBody {
	return s.Body
}

func (s *CloudDeleteTaskResponse) SetHeaders(v map[string]*string) *CloudDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteTaskResponse) SetStatusCode(v int32) *CloudDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteTaskResponse) SetBody(v *CloudDeleteTaskResponseBody) *CloudDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
