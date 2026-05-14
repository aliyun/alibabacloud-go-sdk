// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListEnterprisePausesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListEnterprisePausesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListEnterprisePausesResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListEnterprisePausesResponseBody) *ClinkListEnterprisePausesResponse
	GetBody() *ClinkListEnterprisePausesResponseBody
}

type ClinkListEnterprisePausesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListEnterprisePausesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListEnterprisePausesResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListEnterprisePausesResponse) GoString() string {
	return s.String()
}

func (s *ClinkListEnterprisePausesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListEnterprisePausesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListEnterprisePausesResponse) GetBody() *ClinkListEnterprisePausesResponseBody {
	return s.Body
}

func (s *ClinkListEnterprisePausesResponse) SetHeaders(v map[string]*string) *ClinkListEnterprisePausesResponse {
	s.Headers = v
	return s
}

func (s *ClinkListEnterprisePausesResponse) SetStatusCode(v int32) *ClinkListEnterprisePausesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListEnterprisePausesResponse) SetBody(v *ClinkListEnterprisePausesResponseBody) *ClinkListEnterprisePausesResponse {
	s.Body = v
	return s
}

func (s *ClinkListEnterprisePausesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
