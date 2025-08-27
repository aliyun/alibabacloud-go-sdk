// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderCreateResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderCreateResponseBody) *InsureOrderCreateResponse
	GetBody() *InsureOrderCreateResponseBody
}

type InsureOrderCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderCreateResponse) GetBody() *InsureOrderCreateResponseBody {
	return s.Body
}

func (s *InsureOrderCreateResponse) SetHeaders(v map[string]*string) *InsureOrderCreateResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderCreateResponse) SetStatusCode(v int32) *InsureOrderCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderCreateResponse) SetBody(v *InsureOrderCreateResponseBody) *InsureOrderCreateResponse {
	s.Body = v
	return s
}

func (s *InsureOrderCreateResponse) Validate() error {
	return dara.Validate(s)
}
