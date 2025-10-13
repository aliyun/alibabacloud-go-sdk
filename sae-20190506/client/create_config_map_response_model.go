// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigMapResponseBody) *CreateConfigMapResponse
	GetBody() *CreateConfigMapResponseBody
}

type CreateConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigMapResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigMapResponse) GetBody() *CreateConfigMapResponseBody {
	return s.Body
}

func (s *CreateConfigMapResponse) SetHeaders(v map[string]*string) *CreateConfigMapResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigMapResponse) SetStatusCode(v int32) *CreateConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigMapResponse) SetBody(v *CreateConfigMapResponseBody) *CreateConfigMapResponse {
	s.Body = v
	return s
}

func (s *CreateConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
