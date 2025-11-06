// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailingReasonListForQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFailingReasonListForQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFailingReasonListForQualificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryFailingReasonListForQualificationResponseBody) *QueryFailingReasonListForQualificationResponse
	GetBody() *QueryFailingReasonListForQualificationResponseBody
}

type QueryFailingReasonListForQualificationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFailingReasonListForQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFailingReasonListForQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFailingReasonListForQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFailingReasonListForQualificationResponse) GetBody() *QueryFailingReasonListForQualificationResponseBody {
	return s.Body
}

func (s *QueryFailingReasonListForQualificationResponse) SetHeaders(v map[string]*string) *QueryFailingReasonListForQualificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailingReasonListForQualificationResponse) SetStatusCode(v int32) *QueryFailingReasonListForQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponse) SetBody(v *QueryFailingReasonListForQualificationResponseBody) *QueryFailingReasonListForQualificationResponse {
	s.Body = v
	return s
}

func (s *QueryFailingReasonListForQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
