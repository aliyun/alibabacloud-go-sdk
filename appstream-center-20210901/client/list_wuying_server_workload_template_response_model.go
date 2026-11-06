// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerWorkloadTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWuyingServerWorkloadTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWuyingServerWorkloadTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListWuyingServerWorkloadTemplateResponseBody) *ListWuyingServerWorkloadTemplateResponse
	GetBody() *ListWuyingServerWorkloadTemplateResponseBody
}

type ListWuyingServerWorkloadTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWuyingServerWorkloadTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWuyingServerWorkloadTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerWorkloadTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListWuyingServerWorkloadTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWuyingServerWorkloadTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWuyingServerWorkloadTemplateResponse) GetBody() *ListWuyingServerWorkloadTemplateResponseBody {
	return s.Body
}

func (s *ListWuyingServerWorkloadTemplateResponse) SetHeaders(v map[string]*string) *ListWuyingServerWorkloadTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponse) SetStatusCode(v int32) *ListWuyingServerWorkloadTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponse) SetBody(v *ListWuyingServerWorkloadTemplateResponseBody) *ListWuyingServerWorkloadTemplateResponse {
	s.Body = v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
