// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishGroupTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePublishGroupTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePublishGroupTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatePublishGroupTaskResponseBody) *CreatePublishGroupTaskResponse
	GetBody() *CreatePublishGroupTaskResponseBody
}

type CreatePublishGroupTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublishGroupTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublishGroupTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishGroupTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePublishGroupTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePublishGroupTaskResponse) GetBody() *CreatePublishGroupTaskResponseBody {
	return s.Body
}

func (s *CreatePublishGroupTaskResponse) SetHeaders(v map[string]*string) *CreatePublishGroupTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePublishGroupTaskResponse) SetStatusCode(v int32) *CreatePublishGroupTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublishGroupTaskResponse) SetBody(v *CreatePublishGroupTaskResponseBody) *CreatePublishGroupTaskResponse {
	s.Body = v
	return s
}

func (s *CreatePublishGroupTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
