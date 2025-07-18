// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemoteADBDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRemoteADBDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRemoteADBDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRemoteADBDataSourceResponseBody) *CreateRemoteADBDataSourceResponse
	GetBody() *CreateRemoteADBDataSourceResponseBody
}

type CreateRemoteADBDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRemoteADBDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRemoteADBDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRemoteADBDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateRemoteADBDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRemoteADBDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRemoteADBDataSourceResponse) GetBody() *CreateRemoteADBDataSourceResponseBody {
	return s.Body
}

func (s *CreateRemoteADBDataSourceResponse) SetHeaders(v map[string]*string) *CreateRemoteADBDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateRemoteADBDataSourceResponse) SetStatusCode(v int32) *CreateRemoteADBDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponse) SetBody(v *CreateRemoteADBDataSourceResponseBody) *CreateRemoteADBDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateRemoteADBDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
