// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerAvailableQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryResellerAvailableQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryResellerAvailableQuotaResponse
	GetStatusCode() *int32
	SetBody(v *QueryResellerAvailableQuotaResponseBody) *QueryResellerAvailableQuotaResponse
	GetBody() *QueryResellerAvailableQuotaResponseBody
}

type QueryResellerAvailableQuotaResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResellerAvailableQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResellerAvailableQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerAvailableQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryResellerAvailableQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryResellerAvailableQuotaResponse) GetBody() *QueryResellerAvailableQuotaResponseBody {
	return s.Body
}

func (s *QueryResellerAvailableQuotaResponse) SetHeaders(v map[string]*string) *QueryResellerAvailableQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryResellerAvailableQuotaResponse) SetStatusCode(v int32) *QueryResellerAvailableQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponse) SetBody(v *QueryResellerAvailableQuotaResponseBody) *QueryResellerAvailableQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryResellerAvailableQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
