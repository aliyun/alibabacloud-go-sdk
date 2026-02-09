// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMscpRiskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMscpRiskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMscpRiskInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryMscpRiskInfoResponseBody) *QueryMscpRiskInfoResponse
	GetBody() *QueryMscpRiskInfoResponseBody
}

type QueryMscpRiskInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMscpRiskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMscpRiskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMscpRiskInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMscpRiskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMscpRiskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMscpRiskInfoResponse) GetBody() *QueryMscpRiskInfoResponseBody {
	return s.Body
}

func (s *QueryMscpRiskInfoResponse) SetHeaders(v map[string]*string) *QueryMscpRiskInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMscpRiskInfoResponse) SetStatusCode(v int32) *QueryMscpRiskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMscpRiskInfoResponse) SetBody(v *QueryMscpRiskInfoResponseBody) *QueryMscpRiskInfoResponse {
	s.Body = v
	return s
}

func (s *QueryMscpRiskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
