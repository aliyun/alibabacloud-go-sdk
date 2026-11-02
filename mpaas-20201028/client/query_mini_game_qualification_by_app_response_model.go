// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameQualificationByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMiniGameQualificationByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMiniGameQualificationByAppResponse
	GetStatusCode() *int32
	SetBody(v *QueryMiniGameQualificationByAppResponseBody) *QueryMiniGameQualificationByAppResponse
	GetBody() *QueryMiniGameQualificationByAppResponseBody
}

type QueryMiniGameQualificationByAppResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMiniGameQualificationByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMiniGameQualificationByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppResponse) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMiniGameQualificationByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMiniGameQualificationByAppResponse) GetBody() *QueryMiniGameQualificationByAppResponseBody {
	return s.Body
}

func (s *QueryMiniGameQualificationByAppResponse) SetHeaders(v map[string]*string) *QueryMiniGameQualificationByAppResponse {
	s.Headers = v
	return s
}

func (s *QueryMiniGameQualificationByAppResponse) SetStatusCode(v int32) *QueryMiniGameQualificationByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponse) SetBody(v *QueryMiniGameQualificationByAppResponseBody) *QueryMiniGameQualificationByAppResponse {
	s.Body = v
	return s
}

func (s *QueryMiniGameQualificationByAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
