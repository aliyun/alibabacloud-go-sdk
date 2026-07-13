// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupportedZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySupportedZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySupportedZonesResponse
	GetStatusCode() *int32
	SetBody(v *QuerySupportedZonesResponseBody) *QuerySupportedZonesResponse
	GetBody() *QuerySupportedZonesResponseBody
}

type QuerySupportedZonesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySupportedZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySupportedZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySupportedZonesResponse) GoString() string {
	return s.String()
}

func (s *QuerySupportedZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySupportedZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySupportedZonesResponse) GetBody() *QuerySupportedZonesResponseBody {
	return s.Body
}

func (s *QuerySupportedZonesResponse) SetHeaders(v map[string]*string) *QuerySupportedZonesResponse {
	s.Headers = v
	return s
}

func (s *QuerySupportedZonesResponse) SetStatusCode(v int32) *QuerySupportedZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySupportedZonesResponse) SetBody(v *QuerySupportedZonesResponseBody) *QuerySupportedZonesResponse {
	s.Body = v
	return s
}

func (s *QuerySupportedZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
