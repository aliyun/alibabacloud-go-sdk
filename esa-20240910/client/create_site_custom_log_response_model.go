// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteCustomLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSiteCustomLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSiteCustomLogResponse
	GetStatusCode() *int32
	SetBody(v *CreateSiteCustomLogResponseBody) *CreateSiteCustomLogResponse
	GetBody() *CreateSiteCustomLogResponseBody
}

type CreateSiteCustomLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteCustomLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSiteCustomLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSiteCustomLogResponse) GetBody() *CreateSiteCustomLogResponseBody {
	return s.Body
}

func (s *CreateSiteCustomLogResponse) SetHeaders(v map[string]*string) *CreateSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteCustomLogResponse) SetStatusCode(v int32) *CreateSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteCustomLogResponse) SetBody(v *CreateSiteCustomLogResponseBody) *CreateSiteCustomLogResponse {
	s.Body = v
	return s
}

func (s *CreateSiteCustomLogResponse) Validate() error {
	return dara.Validate(s)
}
