// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudStartTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudStartTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudStartTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudStartTaskResponseBody) *CloudStartTaskResponse
	GetBody() *CloudStartTaskResponseBody
}

type CloudStartTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudStartTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudStartTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudStartTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudStartTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudStartTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudStartTaskResponse) GetBody() *CloudStartTaskResponseBody {
	return s.Body
}

func (s *CloudStartTaskResponse) SetHeaders(v map[string]*string) *CloudStartTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudStartTaskResponse) SetStatusCode(v int32) *CloudStartTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudStartTaskResponse) SetBody(v *CloudStartTaskResponseBody) *CloudStartTaskResponse {
	s.Body = v
	return s
}

func (s *CloudStartTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
