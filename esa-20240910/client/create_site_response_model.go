// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSiteResponse
	GetStatusCode() *int32
	SetBody(v *CreateSiteResponseBody) *CreateSiteResponse
	GetBody() *CreateSiteResponseBody
}

type CreateSiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSiteResponse) GetBody() *CreateSiteResponseBody {
	return s.Body
}

func (s *CreateSiteResponse) SetHeaders(v map[string]*string) *CreateSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteResponse) SetStatusCode(v int32) *CreateSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteResponse) SetBody(v *CreateSiteResponseBody) *CreateSiteResponse {
	s.Body = v
	return s
}

func (s *CreateSiteResponse) Validate() error {
	return dara.Validate(s)
}
