// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenResponseBody) *CreateCenResponse
	GetBody() *CreateCenResponseBody
}

type CreateCenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenResponse) GoString() string {
	return s.String()
}

func (s *CreateCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenResponse) GetBody() *CreateCenResponseBody {
	return s.Body
}

func (s *CreateCenResponse) SetHeaders(v map[string]*string) *CreateCenResponse {
	s.Headers = v
	return s
}

func (s *CreateCenResponse) SetStatusCode(v int32) *CreateCenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenResponse) SetBody(v *CreateCenResponseBody) *CreateCenResponse {
	s.Body = v
	return s
}

func (s *CreateCenResponse) Validate() error {
	return dara.Validate(s)
}
