// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAliasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAliasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAliasesResponse
	GetStatusCode() *int32
	SetBody(v *QueryAliasesResponseBody) *QueryAliasesResponse
	GetBody() *QueryAliasesResponseBody
}

type QueryAliasesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAliasesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAliasesResponse) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAliasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAliasesResponse) GetBody() *QueryAliasesResponseBody {
	return s.Body
}

func (s *QueryAliasesResponse) SetHeaders(v map[string]*string) *QueryAliasesResponse {
	s.Headers = v
	return s
}

func (s *QueryAliasesResponse) SetStatusCode(v int32) *QueryAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAliasesResponse) SetBody(v *QueryAliasesResponseBody) *QueryAliasesResponse {
	s.Body = v
	return s
}

func (s *QueryAliasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
