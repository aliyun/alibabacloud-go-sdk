// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSanityCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSanityCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSanityCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSanityCheckTaskResponseBody) *CreateSanityCheckTaskResponse
	GetBody() *CreateSanityCheckTaskResponseBody
}

type CreateSanityCheckTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSanityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSanityCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSanityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSanityCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSanityCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSanityCheckTaskResponse) GetBody() *CreateSanityCheckTaskResponseBody {
	return s.Body
}

func (s *CreateSanityCheckTaskResponse) SetHeaders(v map[string]*string) *CreateSanityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSanityCheckTaskResponse) SetStatusCode(v int32) *CreateSanityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSanityCheckTaskResponse) SetBody(v *CreateSanityCheckTaskResponseBody) *CreateSanityCheckTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSanityCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
