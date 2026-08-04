// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindPkByHidForLoginWithLegacyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindPkByHidForLoginWithLegacyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindPkByHidForLoginWithLegacyResponse
	GetStatusCode() *int32
	SetBody(v *FindPkByHidForLoginWithLegacyResponseBody) *FindPkByHidForLoginWithLegacyResponse
	GetBody() *FindPkByHidForLoginWithLegacyResponseBody
}

type FindPkByHidForLoginWithLegacyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindPkByHidForLoginWithLegacyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindPkByHidForLoginWithLegacyResponse) String() string {
	return dara.Prettify(s)
}

func (s FindPkByHidForLoginWithLegacyResponse) GoString() string {
	return s.String()
}

func (s *FindPkByHidForLoginWithLegacyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindPkByHidForLoginWithLegacyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindPkByHidForLoginWithLegacyResponse) GetBody() *FindPkByHidForLoginWithLegacyResponseBody {
	return s.Body
}

func (s *FindPkByHidForLoginWithLegacyResponse) SetHeaders(v map[string]*string) *FindPkByHidForLoginWithLegacyResponse {
	s.Headers = v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponse) SetStatusCode(v int32) *FindPkByHidForLoginWithLegacyResponse {
	s.StatusCode = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponse) SetBody(v *FindPkByHidForLoginWithLegacyResponseBody) *FindPkByHidForLoginWithLegacyResponse {
	s.Body = v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
