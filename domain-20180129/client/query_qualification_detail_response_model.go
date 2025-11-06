// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQualificationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryQualificationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryQualificationDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryQualificationDetailResponseBody) *QueryQualificationDetailResponse
	GetBody() *QueryQualificationDetailResponseBody
}

type QueryQualificationDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQualificationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQualificationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryQualificationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryQualificationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryQualificationDetailResponse) GetBody() *QueryQualificationDetailResponseBody {
	return s.Body
}

func (s *QueryQualificationDetailResponse) SetHeaders(v map[string]*string) *QueryQualificationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryQualificationDetailResponse) SetStatusCode(v int32) *QueryQualificationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQualificationDetailResponse) SetBody(v *QueryQualificationDetailResponseBody) *QueryQualificationDetailResponse {
	s.Body = v
	return s
}

func (s *QueryQualificationDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
