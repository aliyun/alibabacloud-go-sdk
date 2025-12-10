// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteResourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PromoteResourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PromoteResourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *PromoteResourceAccountResponseBody) *PromoteResourceAccountResponse
	GetBody() *PromoteResourceAccountResponseBody
}

type PromoteResourceAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PromoteResourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s PromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PromoteResourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PromoteResourceAccountResponse) GetBody() *PromoteResourceAccountResponseBody {
	return s.Body
}

func (s *PromoteResourceAccountResponse) SetHeaders(v map[string]*string) *PromoteResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *PromoteResourceAccountResponse) SetStatusCode(v int32) *PromoteResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PromoteResourceAccountResponse) SetBody(v *PromoteResourceAccountResponseBody) *PromoteResourceAccountResponse {
	s.Body = v
	return s
}

func (s *PromoteResourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
