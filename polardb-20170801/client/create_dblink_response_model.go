// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBLinkResponseBody) *CreateDBLinkResponse
	GetBody() *CreateDBLinkResponseBody
}

type CreateDBLinkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateDBLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBLinkResponse) GetBody() *CreateDBLinkResponseBody {
	return s.Body
}

func (s *CreateDBLinkResponse) SetHeaders(v map[string]*string) *CreateDBLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateDBLinkResponse) SetStatusCode(v int32) *CreateDBLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBLinkResponse) SetBody(v *CreateDBLinkResponseBody) *CreateDBLinkResponse {
	s.Body = v
	return s
}

func (s *CreateDBLinkResponse) Validate() error {
	return dara.Validate(s)
}
