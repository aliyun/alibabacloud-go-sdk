// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateResponse
	GetStatusCode() *int32
	SetBody(v *AssociateResponseBody) *AssociateResponse
	GetBody() *AssociateResponseBody
}

type AssociateResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateResponse) GoString() string {
	return s.String()
}

func (s *AssociateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateResponse) GetBody() *AssociateResponseBody {
	return s.Body
}

func (s *AssociateResponse) SetHeaders(v map[string]*string) *AssociateResponse {
	s.Headers = v
	return s
}

func (s *AssociateResponse) SetStatusCode(v int32) *AssociateResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResponse) SetBody(v *AssociateResponseBody) *AssociateResponse {
	s.Body = v
	return s
}

func (s *AssociateResponse) Validate() error {
	return dara.Validate(s)
}
