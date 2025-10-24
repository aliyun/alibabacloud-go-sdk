// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTemplateListResponseBody) *QueryTemplateListResponse
	GetBody() *QueryTemplateListResponseBody
}

type QueryTemplateListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTemplateListResponse) GetBody() *QueryTemplateListResponseBody {
	return s.Body
}

func (s *QueryTemplateListResponse) SetHeaders(v map[string]*string) *QueryTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateListResponse) SetStatusCode(v int32) *QueryTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTemplateListResponse) SetBody(v *QueryTemplateListResponseBody) *QueryTemplateListResponse {
	s.Body = v
	return s
}

func (s *QueryTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
