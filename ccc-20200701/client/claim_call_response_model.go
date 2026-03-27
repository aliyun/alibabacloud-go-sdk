// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClaimCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClaimCallResponse
	GetStatusCode() *int32
	SetBody(v *ClaimCallResponseBody) *ClaimCallResponse
	GetBody() *ClaimCallResponseBody
}

type ClaimCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClaimCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClaimCallResponse) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallResponse) GoString() string {
	return s.String()
}

func (s *ClaimCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClaimCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClaimCallResponse) GetBody() *ClaimCallResponseBody {
	return s.Body
}

func (s *ClaimCallResponse) SetHeaders(v map[string]*string) *ClaimCallResponse {
	s.Headers = v
	return s
}

func (s *ClaimCallResponse) SetStatusCode(v int32) *ClaimCallResponse {
	s.StatusCode = &v
	return s
}

func (s *ClaimCallResponse) SetBody(v *ClaimCallResponseBody) *ClaimCallResponse {
	s.Body = v
	return s
}

func (s *ClaimCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
