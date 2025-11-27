// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonCodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAddonCodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAddonCodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAddonCodeTemplateResponseBody) *GetAddonCodeTemplateResponse
	GetBody() *GetAddonCodeTemplateResponseBody
}

type GetAddonCodeTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddonCodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddonCodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAddonCodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAddonCodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAddonCodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAddonCodeTemplateResponse) GetBody() *GetAddonCodeTemplateResponseBody {
	return s.Body
}

func (s *GetAddonCodeTemplateResponse) SetHeaders(v map[string]*string) *GetAddonCodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAddonCodeTemplateResponse) SetStatusCode(v int32) *GetAddonCodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddonCodeTemplateResponse) SetBody(v *GetAddonCodeTemplateResponseBody) *GetAddonCodeTemplateResponse {
	s.Body = v
	return s
}

func (s *GetAddonCodeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
