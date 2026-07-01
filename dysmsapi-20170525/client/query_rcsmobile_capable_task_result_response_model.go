// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRCSMobileCapableTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRCSMobileCapableTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryRCSMobileCapableTaskResultResponseBody) *QueryRCSMobileCapableTaskResultResponse
	GetBody() *QueryRCSMobileCapableTaskResultResponseBody
}

type QueryRCSMobileCapableTaskResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRCSMobileCapableTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRCSMobileCapableTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableTaskResultResponse) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRCSMobileCapableTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRCSMobileCapableTaskResultResponse) GetBody() *QueryRCSMobileCapableTaskResultResponseBody {
	return s.Body
}

func (s *QueryRCSMobileCapableTaskResultResponse) SetHeaders(v map[string]*string) *QueryRCSMobileCapableTaskResultResponse {
	s.Headers = v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponse) SetStatusCode(v int32) *QueryRCSMobileCapableTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponse) SetBody(v *QueryRCSMobileCapableTaskResultResponseBody) *QueryRCSMobileCapableTaskResultResponse {
	s.Body = v
	return s
}

func (s *QueryRCSMobileCapableTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
