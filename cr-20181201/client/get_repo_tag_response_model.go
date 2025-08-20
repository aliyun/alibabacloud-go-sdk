// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoTagResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoTagResponseBody) *GetRepoTagResponse
	GetBody() *GetRepoTagResponseBody
}

type GetRepoTagResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoTagResponse) GetBody() *GetRepoTagResponseBody {
	return s.Body
}

func (s *GetRepoTagResponse) SetHeaders(v map[string]*string) *GetRepoTagResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagResponse) SetStatusCode(v int32) *GetRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagResponse) SetBody(v *GetRepoTagResponseBody) *GetRepoTagResponse {
	s.Body = v
	return s
}

func (s *GetRepoTagResponse) Validate() error {
	return dara.Validate(s)
}
