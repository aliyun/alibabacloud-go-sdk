// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoCreationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoCreationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoCreationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoCreationTaskResponseBody) *CreateVideoCreationTaskResponse
	GetBody() *CreateVideoCreationTaskResponseBody
}

type CreateVideoCreationTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoCreationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoCreationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoCreationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoCreationTaskResponse) GetBody() *CreateVideoCreationTaskResponseBody {
	return s.Body
}

func (s *CreateVideoCreationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoCreationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoCreationTaskResponse) SetStatusCode(v int32) *CreateVideoCreationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoCreationTaskResponse) SetBody(v *CreateVideoCreationTaskResponseBody) *CreateVideoCreationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoCreationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
