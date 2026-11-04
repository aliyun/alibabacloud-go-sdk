// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTimeTemplateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTimeTemplateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTimeTemplateDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryTimeTemplateDetailResponseBody) *QueryTimeTemplateDetailResponse
	GetBody() *QueryTimeTemplateDetailResponseBody
}

type QueryTimeTemplateDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTimeTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTimeTemplateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTimeTemplateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTimeTemplateDetailResponse) GetBody() *QueryTimeTemplateDetailResponseBody {
	return s.Body
}

func (s *QueryTimeTemplateDetailResponse) SetHeaders(v map[string]*string) *QueryTimeTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTimeTemplateDetailResponse) SetStatusCode(v int32) *QueryTimeTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTimeTemplateDetailResponse) SetBody(v *QueryTimeTemplateDetailResponseBody) *QueryTimeTemplateDetailResponse {
	s.Body = v
	return s
}

func (s *QueryTimeTemplateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
