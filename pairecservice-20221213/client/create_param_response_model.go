// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateParamResponse
	GetStatusCode() *int32
	SetBody(v *CreateParamResponseBody) *CreateParamResponse
	GetBody() *CreateParamResponseBody
}

type CreateParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateParamResponse) GoString() string {
	return s.String()
}

func (s *CreateParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateParamResponse) GetBody() *CreateParamResponseBody {
	return s.Body
}

func (s *CreateParamResponse) SetHeaders(v map[string]*string) *CreateParamResponse {
	s.Headers = v
	return s
}

func (s *CreateParamResponse) SetStatusCode(v int32) *CreateParamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParamResponse) SetBody(v *CreateParamResponseBody) *CreateParamResponse {
	s.Body = v
	return s
}

func (s *CreateParamResponse) Validate() error {
	return dara.Validate(s)
}
