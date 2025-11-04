// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateMaterialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateMaterialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateMaterialsResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateMaterialsResponseBody) *GetTemplateMaterialsResponse
	GetBody() *GetTemplateMaterialsResponseBody
}

type GetTemplateMaterialsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateMaterialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateMaterialsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateMaterialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateMaterialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateMaterialsResponse) GetBody() *GetTemplateMaterialsResponseBody {
	return s.Body
}

func (s *GetTemplateMaterialsResponse) SetHeaders(v map[string]*string) *GetTemplateMaterialsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateMaterialsResponse) SetStatusCode(v int32) *GetTemplateMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateMaterialsResponse) SetBody(v *GetTemplateMaterialsResponseBody) *GetTemplateMaterialsResponse {
	s.Body = v
	return s
}

func (s *GetTemplateMaterialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
