// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEmasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenEmasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenEmasServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenEmasServiceResponseBody) *OpenEmasServiceResponse
	GetBody() *OpenEmasServiceResponseBody
}

type OpenEmasServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenEmasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenEmasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenEmasServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenEmasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenEmasServiceResponse) GetBody() *OpenEmasServiceResponseBody {
	return s.Body
}

func (s *OpenEmasServiceResponse) SetHeaders(v map[string]*string) *OpenEmasServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenEmasServiceResponse) SetStatusCode(v int32) *OpenEmasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenEmasServiceResponse) SetBody(v *OpenEmasServiceResponseBody) *OpenEmasServiceResponse {
	s.Body = v
	return s
}

func (s *OpenEmasServiceResponse) Validate() error {
	return dara.Validate(s)
}
