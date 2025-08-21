// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateWebCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateWebCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateWebCertResponse
	GetStatusCode() *int32
	SetBody(v *AssociateWebCertResponseBody) *AssociateWebCertResponse
	GetBody() *AssociateWebCertResponseBody
}

type AssociateWebCertResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateWebCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateWebCertResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateWebCertResponse) GoString() string {
	return s.String()
}

func (s *AssociateWebCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateWebCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateWebCertResponse) GetBody() *AssociateWebCertResponseBody {
	return s.Body
}

func (s *AssociateWebCertResponse) SetHeaders(v map[string]*string) *AssociateWebCertResponse {
	s.Headers = v
	return s
}

func (s *AssociateWebCertResponse) SetStatusCode(v int32) *AssociateWebCertResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateWebCertResponse) SetBody(v *AssociateWebCertResponseBody) *AssociateWebCertResponse {
	s.Body = v
	return s
}

func (s *AssociateWebCertResponse) Validate() error {
	return dara.Validate(s)
}
