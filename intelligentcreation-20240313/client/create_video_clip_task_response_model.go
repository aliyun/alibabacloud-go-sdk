// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoClipTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoClipTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoClipTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoClipTaskResponseBody) *CreateVideoClipTaskResponse
	GetBody() *CreateVideoClipTaskResponseBody
}

type CreateVideoClipTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoClipTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoClipTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoClipTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoClipTaskResponse) GetBody() *CreateVideoClipTaskResponseBody {
	return s.Body
}

func (s *CreateVideoClipTaskResponse) SetHeaders(v map[string]*string) *CreateVideoClipTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoClipTaskResponse) SetStatusCode(v int32) *CreateVideoClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoClipTaskResponse) SetBody(v *CreateVideoClipTaskResponseBody) *CreateVideoClipTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoClipTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
