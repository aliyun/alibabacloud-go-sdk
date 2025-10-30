// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivateAccessTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivateAccessTagResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivateAccessTagResponseBody) *CreatePrivateAccessTagResponse
	GetBody() *CreatePrivateAccessTagResponseBody
}

type CreatePrivateAccessTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivateAccessTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivateAccessTagResponse) GetBody() *CreatePrivateAccessTagResponseBody {
	return s.Body
}

func (s *CreatePrivateAccessTagResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessTagResponse) SetStatusCode(v int32) *CreatePrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessTagResponse) SetBody(v *CreatePrivateAccessTagResponseBody) *CreatePrivateAccessTagResponse {
	s.Body = v
	return s
}

func (s *CreatePrivateAccessTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
