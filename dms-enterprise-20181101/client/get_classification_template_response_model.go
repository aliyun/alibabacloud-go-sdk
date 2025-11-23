// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClassificationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClassificationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClassificationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetClassificationTemplateResponseBody) *GetClassificationTemplateResponse
	GetBody() *GetClassificationTemplateResponseBody
}

type GetClassificationTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClassificationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClassificationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClassificationTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetClassificationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClassificationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClassificationTemplateResponse) GetBody() *GetClassificationTemplateResponseBody {
	return s.Body
}

func (s *GetClassificationTemplateResponse) SetHeaders(v map[string]*string) *GetClassificationTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetClassificationTemplateResponse) SetStatusCode(v int32) *GetClassificationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClassificationTemplateResponse) SetBody(v *GetClassificationTemplateResponseBody) *GetClassificationTemplateResponse {
	s.Body = v
	return s
}

func (s *GetClassificationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
