// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEaiEcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEaiEcsResponse
	GetStatusCode() *int32
	SetBody(v *CreateEaiEcsResponseBody) *CreateEaiEcsResponse
	GetBody() *CreateEaiEcsResponseBody
}

type CreateEaiEcsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiEcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiEcsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEaiEcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEaiEcsResponse) GetBody() *CreateEaiEcsResponseBody {
	return s.Body
}

func (s *CreateEaiEcsResponse) SetHeaders(v map[string]*string) *CreateEaiEcsResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiEcsResponse) SetStatusCode(v int32) *CreateEaiEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiEcsResponse) SetBody(v *CreateEaiEcsResponseBody) *CreateEaiEcsResponse {
	s.Body = v
	return s
}

func (s *CreateEaiEcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
