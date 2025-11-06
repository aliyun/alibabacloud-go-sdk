// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMdsCubeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMdsCubeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMdsCubeTaskResponseBody) *CreateMdsCubeTaskResponse
	GetBody() *CreateMdsCubeTaskResponseBody
}

type CreateMdsCubeTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMdsCubeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMdsCubeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMdsCubeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMdsCubeTaskResponse) GetBody() *CreateMdsCubeTaskResponseBody {
	return s.Body
}

func (s *CreateMdsCubeTaskResponse) SetHeaders(v map[string]*string) *CreateMdsCubeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMdsCubeTaskResponse) SetStatusCode(v int32) *CreateMdsCubeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMdsCubeTaskResponse) SetBody(v *CreateMdsCubeTaskResponseBody) *CreateMdsCubeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMdsCubeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
