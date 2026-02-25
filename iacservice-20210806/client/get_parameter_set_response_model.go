// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParameterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParameterSetResponse
	GetStatusCode() *int32
	SetBody(v *GetParameterSetResponseBody) *GetParameterSetResponse
	GetBody() *GetParameterSetResponseBody
}

type GetParameterSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParameterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetResponse) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParameterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParameterSetResponse) GetBody() *GetParameterSetResponseBody {
	return s.Body
}

func (s *GetParameterSetResponse) SetHeaders(v map[string]*string) *GetParameterSetResponse {
	s.Headers = v
	return s
}

func (s *GetParameterSetResponse) SetStatusCode(v int32) *GetParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParameterSetResponse) SetBody(v *GetParameterSetResponseBody) *GetParameterSetResponse {
	s.Body = v
	return s
}

func (s *GetParameterSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
