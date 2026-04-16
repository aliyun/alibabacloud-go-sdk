// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateTaskResponseBody) *CloudCreateTaskResponse
	GetBody() *CloudCreateTaskResponseBody
}

type CloudCreateTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateTaskResponse) GetBody() *CloudCreateTaskResponseBody {
	return s.Body
}

func (s *CloudCreateTaskResponse) SetHeaders(v map[string]*string) *CloudCreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateTaskResponse) SetStatusCode(v int32) *CloudCreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateTaskResponse) SetBody(v *CloudCreateTaskResponseBody) *CloudCreateTaskResponse {
	s.Body = v
	return s
}

func (s *CloudCreateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
