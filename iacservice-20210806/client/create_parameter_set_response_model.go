// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateParameterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateParameterSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateParameterSetResponseBody) *CreateParameterSetResponse
	GetBody() *CreateParameterSetResponseBody
}

type CreateParameterSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParameterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateParameterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateParameterSetResponse) GetBody() *CreateParameterSetResponseBody {
	return s.Body
}

func (s *CreateParameterSetResponse) SetHeaders(v map[string]*string) *CreateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterSetResponse) SetStatusCode(v int32) *CreateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParameterSetResponse) SetBody(v *CreateParameterSetResponseBody) *CreateParameterSetResponse {
	s.Body = v
	return s
}

func (s *CreateParameterSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
