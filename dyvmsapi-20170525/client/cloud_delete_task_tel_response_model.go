// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskTelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteTaskTelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteTaskTelResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteTaskTelResponseBody) *CloudDeleteTaskTelResponse
	GetBody() *CloudDeleteTaskTelResponseBody
}

type CloudDeleteTaskTelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteTaskTelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteTaskTelResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskTelResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskTelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteTaskTelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteTaskTelResponse) GetBody() *CloudDeleteTaskTelResponseBody {
	return s.Body
}

func (s *CloudDeleteTaskTelResponse) SetHeaders(v map[string]*string) *CloudDeleteTaskTelResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteTaskTelResponse) SetStatusCode(v int32) *CloudDeleteTaskTelResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteTaskTelResponse) SetBody(v *CloudDeleteTaskTelResponseBody) *CloudDeleteTaskTelResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteTaskTelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
