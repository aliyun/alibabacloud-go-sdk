// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDnsGtmCnameRrCanUseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateDnsGtmCnameRrCanUseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateDnsGtmCnameRrCanUseResponse
	GetStatusCode() *int32
	SetBody(v *ValidateDnsGtmCnameRrCanUseResponseBody) *ValidateDnsGtmCnameRrCanUseResponse
	GetBody() *ValidateDnsGtmCnameRrCanUseResponseBody
}

type ValidateDnsGtmCnameRrCanUseResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateDnsGtmCnameRrCanUseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDnsGtmCnameRrCanUseResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateDnsGtmCnameRrCanUseResponse) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) GetBody() *ValidateDnsGtmCnameRrCanUseResponseBody {
	return s.Body
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) SetHeaders(v map[string]*string) *ValidateDnsGtmCnameRrCanUseResponse {
	s.Headers = v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) SetStatusCode(v int32) *ValidateDnsGtmCnameRrCanUseResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) SetBody(v *ValidateDnsGtmCnameRrCanUseResponseBody) *ValidateDnsGtmCnameRrCanUseResponse {
	s.Body = v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseResponse) Validate() error {
	return dara.Validate(s)
}
