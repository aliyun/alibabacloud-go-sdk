// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryEnterpriseInfoResponseBody) *QueryEnterpriseInfoResponse
	GetBody() *QueryEnterpriseInfoResponseBody
}

type QueryEnterpriseInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEnterpriseInfoResponse) GetBody() *QueryEnterpriseInfoResponseBody {
	return s.Body
}

func (s *QueryEnterpriseInfoResponse) SetHeaders(v map[string]*string) *QueryEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEnterpriseInfoResponse) SetStatusCode(v int32) *QueryEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnterpriseInfoResponse) SetBody(v *QueryEnterpriseInfoResponseBody) *QueryEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *QueryEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
