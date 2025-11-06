// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsMiniprogramTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMdsMiniprogramTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMdsMiniprogramTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMdsMiniprogramTaskResponseBody) *CreateMdsMiniprogramTaskResponse
	GetBody() *CreateMdsMiniprogramTaskResponseBody
}

type CreateMdsMiniprogramTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMdsMiniprogramTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMdsMiniprogramTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsMiniprogramTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMdsMiniprogramTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMdsMiniprogramTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMdsMiniprogramTaskResponse) GetBody() *CreateMdsMiniprogramTaskResponseBody {
	return s.Body
}

func (s *CreateMdsMiniprogramTaskResponse) SetHeaders(v map[string]*string) *CreateMdsMiniprogramTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMdsMiniprogramTaskResponse) SetStatusCode(v int32) *CreateMdsMiniprogramTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponse) SetBody(v *CreateMdsMiniprogramTaskResponseBody) *CreateMdsMiniprogramTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMdsMiniprogramTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
