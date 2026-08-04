// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecurityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecurityInfoResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecurityInfoResponseBody) *QuerySecurityInfoResponse
	GetBody() *QuerySecurityInfoResponseBody
}

type QuerySecurityInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecurityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySecurityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecurityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecurityInfoResponse) GetBody() *QuerySecurityInfoResponseBody {
	return s.Body
}

func (s *QuerySecurityInfoResponse) SetHeaders(v map[string]*string) *QuerySecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySecurityInfoResponse) SetStatusCode(v int32) *QuerySecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecurityInfoResponse) SetBody(v *QuerySecurityInfoResponseBody) *QuerySecurityInfoResponse {
	s.Body = v
	return s
}

func (s *QuerySecurityInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
