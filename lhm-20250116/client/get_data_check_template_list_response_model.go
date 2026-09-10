// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckTemplateListResponseBody) *GetDataCheckTemplateListResponse
	GetBody() *GetDataCheckTemplateListResponseBody
}

type GetDataCheckTemplateListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckTemplateListResponse) GetBody() *GetDataCheckTemplateListResponseBody {
	return s.Body
}

func (s *GetDataCheckTemplateListResponse) SetHeaders(v map[string]*string) *GetDataCheckTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckTemplateListResponse) SetStatusCode(v int32) *GetDataCheckTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckTemplateListResponse) SetBody(v *GetDataCheckTemplateListResponseBody) *GetDataCheckTemplateListResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
