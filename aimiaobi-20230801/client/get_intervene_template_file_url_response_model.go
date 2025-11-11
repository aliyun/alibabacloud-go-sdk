// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneTemplateFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterveneTemplateFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterveneTemplateFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetInterveneTemplateFileUrlResponseBody) *GetInterveneTemplateFileUrlResponse
	GetBody() *GetInterveneTemplateFileUrlResponseBody
}

type GetInterveneTemplateFileUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneTemplateFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterveneTemplateFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterveneTemplateFileUrlResponse) GetBody() *GetInterveneTemplateFileUrlResponseBody {
	return s.Body
}

func (s *GetInterveneTemplateFileUrlResponse) SetHeaders(v map[string]*string) *GetInterveneTemplateFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetStatusCode(v int32) *GetInterveneTemplateFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetBody(v *GetInterveneTemplateFileUrlResponseBody) *GetInterveneTemplateFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
