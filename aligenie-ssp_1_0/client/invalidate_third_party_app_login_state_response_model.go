// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateThirdPartyAppLoginStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvalidateThirdPartyAppLoginStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvalidateThirdPartyAppLoginStateResponse
	GetStatusCode() *int32
	SetBody(v *InvalidateThirdPartyAppLoginStateResponseBody) *InvalidateThirdPartyAppLoginStateResponse
	GetBody() *InvalidateThirdPartyAppLoginStateResponseBody
}

type InvalidateThirdPartyAppLoginStateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidateThirdPartyAppLoginStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateResponse) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateResponse) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvalidateThirdPartyAppLoginStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvalidateThirdPartyAppLoginStateResponse) GetBody() *InvalidateThirdPartyAppLoginStateResponseBody {
	return s.Body
}

func (s *InvalidateThirdPartyAppLoginStateResponse) SetHeaders(v map[string]*string) *InvalidateThirdPartyAppLoginStateResponse {
	s.Headers = v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponse) SetStatusCode(v int32) *InvalidateThirdPartyAppLoginStateResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponse) SetBody(v *InvalidateThirdPartyAppLoginStateResponseBody) *InvalidateThirdPartyAppLoginStateResponse {
	s.Body = v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponse) Validate() error {
	return dara.Validate(s)
}
