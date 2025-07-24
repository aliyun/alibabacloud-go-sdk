// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateParameterConstraintsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse
	GetBody() *GetServiceTemplateParameterConstraintsResponseBody
}

type GetServiceTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceTemplateParameterConstraintsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceTemplateParameterConstraintsResponse) GetBody() *GetServiceTemplateParameterConstraintsResponseBody {
	return s.Body
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) Validate() error {
	return dara.Validate(s)
}
