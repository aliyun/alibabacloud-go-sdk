// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteNameExclusiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteNameExclusiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteNameExclusiveResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteNameExclusiveResponseBody) *UpdateSiteNameExclusiveResponse
	GetBody() *UpdateSiteNameExclusiveResponseBody
}

type UpdateSiteNameExclusiveResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteNameExclusiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteNameExclusiveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteNameExclusiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteNameExclusiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteNameExclusiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteNameExclusiveResponse) GetBody() *UpdateSiteNameExclusiveResponseBody {
	return s.Body
}

func (s *UpdateSiteNameExclusiveResponse) SetHeaders(v map[string]*string) *UpdateSiteNameExclusiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteNameExclusiveResponse) SetStatusCode(v int32) *UpdateSiteNameExclusiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteNameExclusiveResponse) SetBody(v *UpdateSiteNameExclusiveResponseBody) *UpdateSiteNameExclusiveResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteNameExclusiveResponse) Validate() error {
	return dara.Validate(s)
}
