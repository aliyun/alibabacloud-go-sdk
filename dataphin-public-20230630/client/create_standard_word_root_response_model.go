// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardWordRootResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardWordRootResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardWordRootResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardWordRootResponseBody) *CreateStandardWordRootResponse
	GetBody() *CreateStandardWordRootResponseBody
}

type CreateStandardWordRootResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardWordRootResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardWordRootResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardWordRootResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardWordRootResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardWordRootResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardWordRootResponse) GetBody() *CreateStandardWordRootResponseBody {
	return s.Body
}

func (s *CreateStandardWordRootResponse) SetHeaders(v map[string]*string) *CreateStandardWordRootResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardWordRootResponse) SetStatusCode(v int32) *CreateStandardWordRootResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardWordRootResponse) SetBody(v *CreateStandardWordRootResponseBody) *CreateStandardWordRootResponse {
	s.Body = v
	return s
}

func (s *CreateStandardWordRootResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
