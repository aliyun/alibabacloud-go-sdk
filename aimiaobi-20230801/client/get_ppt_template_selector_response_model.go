// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptTemplateSelectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPptTemplateSelectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPptTemplateSelectorResponse
	GetStatusCode() *int32
	SetBody(v *GetPptTemplateSelectorResponseBody) *GetPptTemplateSelectorResponse
	GetBody() *GetPptTemplateSelectorResponseBody
}

type GetPptTemplateSelectorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPptTemplateSelectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPptTemplateSelectorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorResponse) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPptTemplateSelectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPptTemplateSelectorResponse) GetBody() *GetPptTemplateSelectorResponseBody {
	return s.Body
}

func (s *GetPptTemplateSelectorResponse) SetHeaders(v map[string]*string) *GetPptTemplateSelectorResponse {
	s.Headers = v
	return s
}

func (s *GetPptTemplateSelectorResponse) SetStatusCode(v int32) *GetPptTemplateSelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPptTemplateSelectorResponse) SetBody(v *GetPptTemplateSelectorResponseBody) *GetPptTemplateSelectorResponse {
	s.Body = v
	return s
}

func (s *GetPptTemplateSelectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
