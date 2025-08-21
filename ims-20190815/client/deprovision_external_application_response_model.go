// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionExternalApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeprovisionExternalApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeprovisionExternalApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeprovisionExternalApplicationResponseBody) *DeprovisionExternalApplicationResponse
	GetBody() *DeprovisionExternalApplicationResponseBody
}

type DeprovisionExternalApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionExternalApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeprovisionExternalApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeprovisionExternalApplicationResponse) GetBody() *DeprovisionExternalApplicationResponseBody {
	return s.Body
}

func (s *DeprovisionExternalApplicationResponse) SetHeaders(v map[string]*string) *DeprovisionExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionExternalApplicationResponse) SetStatusCode(v int32) *DeprovisionExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionExternalApplicationResponse) SetBody(v *DeprovisionExternalApplicationResponseBody) *DeprovisionExternalApplicationResponse {
	s.Body = v
	return s
}

func (s *DeprovisionExternalApplicationResponse) Validate() error {
	return dara.Validate(s)
}
