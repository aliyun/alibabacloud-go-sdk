// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmarttagTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmarttagTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmarttagTemplateListResponseBody) *QuerySmarttagTemplateListResponse
	GetBody() *QuerySmarttagTemplateListResponseBody
}

type QuerySmarttagTemplateListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmarttagTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmarttagTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QuerySmarttagTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmarttagTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmarttagTemplateListResponse) GetBody() *QuerySmarttagTemplateListResponseBody {
	return s.Body
}

func (s *QuerySmarttagTemplateListResponse) SetHeaders(v map[string]*string) *QuerySmarttagTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QuerySmarttagTemplateListResponse) SetStatusCode(v int32) *QuerySmarttagTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmarttagTemplateListResponse) SetBody(v *QuerySmarttagTemplateListResponseBody) *QuerySmarttagTemplateListResponse {
	s.Body = v
	return s
}

func (s *QuerySmarttagTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
