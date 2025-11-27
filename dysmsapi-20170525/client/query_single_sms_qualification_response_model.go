// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleSmsQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySingleSmsQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySingleSmsQualificationResponse
	GetStatusCode() *int32
	SetBody(v *QuerySingleSmsQualificationResponseBody) *QuerySingleSmsQualificationResponse
	GetBody() *QuerySingleSmsQualificationResponseBody
}

type QuerySingleSmsQualificationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySingleSmsQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySingleSmsQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySingleSmsQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySingleSmsQualificationResponse) GetBody() *QuerySingleSmsQualificationResponseBody {
	return s.Body
}

func (s *QuerySingleSmsQualificationResponse) SetHeaders(v map[string]*string) *QuerySingleSmsQualificationResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleSmsQualificationResponse) SetStatusCode(v int32) *QuerySingleSmsQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySingleSmsQualificationResponse) SetBody(v *QuerySingleSmsQualificationResponseBody) *QuerySingleSmsQualificationResponse {
	s.Body = v
	return s
}

func (s *QuerySingleSmsQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
