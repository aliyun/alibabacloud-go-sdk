// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteAccessTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteAccessTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteAccessTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteAccessTypeResponseBody) *UpdateSiteAccessTypeResponse
	GetBody() *UpdateSiteAccessTypeResponseBody
}

type UpdateSiteAccessTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteAccessTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteAccessTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteAccessTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteAccessTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteAccessTypeResponse) GetBody() *UpdateSiteAccessTypeResponseBody {
	return s.Body
}

func (s *UpdateSiteAccessTypeResponse) SetHeaders(v map[string]*string) *UpdateSiteAccessTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteAccessTypeResponse) SetStatusCode(v int32) *UpdateSiteAccessTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteAccessTypeResponse) SetBody(v *UpdateSiteAccessTypeResponseBody) *UpdateSiteAccessTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteAccessTypeResponse) Validate() error {
	return dara.Validate(s)
}
