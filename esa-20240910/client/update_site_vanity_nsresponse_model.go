// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteVanityNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteVanityNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteVanityNSResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteVanityNSResponseBody) *UpdateSiteVanityNSResponse
	GetBody() *UpdateSiteVanityNSResponseBody
}

type UpdateSiteVanityNSResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteVanityNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteVanityNSResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteVanityNSResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteVanityNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteVanityNSResponse) GetBody() *UpdateSiteVanityNSResponseBody {
	return s.Body
}

func (s *UpdateSiteVanityNSResponse) SetHeaders(v map[string]*string) *UpdateSiteVanityNSResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteVanityNSResponse) SetStatusCode(v int32) *UpdateSiteVanityNSResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteVanityNSResponse) SetBody(v *UpdateSiteVanityNSResponseBody) *UpdateSiteVanityNSResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteVanityNSResponse) Validate() error {
	return dara.Validate(s)
}
