// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoTagResponseBody) *CreateRepoTagResponse
	GetBody() *CreateRepoTagResponseBody
}

type CreateRepoTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoTagResponse) GetBody() *CreateRepoTagResponseBody {
	return s.Body
}

func (s *CreateRepoTagResponse) SetHeaders(v map[string]*string) *CreateRepoTagResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTagResponse) SetStatusCode(v int32) *CreateRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTagResponse) SetBody(v *CreateRepoTagResponseBody) *CreateRepoTagResponse {
	s.Body = v
	return s
}

func (s *CreateRepoTagResponse) Validate() error {
	return dara.Validate(s)
}
