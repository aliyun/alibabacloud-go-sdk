// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageTaskResponseBody) *CreateImageTaskResponse
	GetBody() *CreateImageTaskResponseBody
}

type CreateImageTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageTaskResponse) GetBody() *CreateImageTaskResponseBody {
	return s.Body
}

func (s *CreateImageTaskResponse) SetHeaders(v map[string]*string) *CreateImageTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageTaskResponse) SetStatusCode(v int32) *CreateImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageTaskResponse) SetBody(v *CreateImageTaskResponseBody) *CreateImageTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
