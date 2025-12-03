// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEciResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEaiEciResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEaiEciResponse
	GetStatusCode() *int32
	SetBody(v *CreateEaiEciResponseBody) *CreateEaiEciResponse
	GetBody() *CreateEaiEciResponseBody
}

type CreateEaiEciResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiEciResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiEciResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiEciResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEaiEciResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEaiEciResponse) GetBody() *CreateEaiEciResponseBody {
	return s.Body
}

func (s *CreateEaiEciResponse) SetHeaders(v map[string]*string) *CreateEaiEciResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiEciResponse) SetStatusCode(v int32) *CreateEaiEciResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiEciResponse) SetBody(v *CreateEaiEciResponseBody) *CreateEaiEciResponse {
	s.Body = v
	return s
}

func (s *CreateEaiEciResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
