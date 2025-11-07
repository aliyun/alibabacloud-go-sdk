// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParameterConstraintsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateParameterConstraintsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateParameterConstraintsResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateParameterConstraintsResponseBody) *GetTemplateParameterConstraintsResponse
	GetBody() *GetTemplateParameterConstraintsResponseBody
}

type GetTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateParameterConstraintsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateParameterConstraintsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateParameterConstraintsResponse) GetBody() *GetTemplateParameterConstraintsResponseBody {
	return s.Body
}

func (s *GetTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponse) SetBody(v *GetTemplateParameterConstraintsResponseBody) *GetTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

func (s *GetTemplateParameterConstraintsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
