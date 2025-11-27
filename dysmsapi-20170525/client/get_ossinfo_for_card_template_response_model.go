// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSInfoForCardTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOSSInfoForCardTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOSSInfoForCardTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetOSSInfoForCardTemplateResponseBody) *GetOSSInfoForCardTemplateResponse
	GetBody() *GetOSSInfoForCardTemplateResponseBody
}

type GetOSSInfoForCardTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSInfoForCardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSInfoForCardTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOSSInfoForCardTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOSSInfoForCardTemplateResponse) GetBody() *GetOSSInfoForCardTemplateResponseBody {
	return s.Body
}

func (s *GetOSSInfoForCardTemplateResponse) SetHeaders(v map[string]*string) *GetOSSInfoForCardTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetOSSInfoForCardTemplateResponse) SetStatusCode(v int32) *GetOSSInfoForCardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponse) SetBody(v *GetOSSInfoForCardTemplateResponseBody) *GetOSSInfoForCardTemplateResponse {
	s.Body = v
	return s
}

func (s *GetOSSInfoForCardTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
