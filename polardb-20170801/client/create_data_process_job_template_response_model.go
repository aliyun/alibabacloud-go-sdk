// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataProcessJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataProcessJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataProcessJobTemplateResponseBody) *CreateDataProcessJobTemplateResponse
	GetBody() *CreateDataProcessJobTemplateResponseBody
}

type CreateDataProcessJobTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataProcessJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataProcessJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataProcessJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataProcessJobTemplateResponse) GetBody() *CreateDataProcessJobTemplateResponseBody {
	return s.Body
}

func (s *CreateDataProcessJobTemplateResponse) SetHeaders(v map[string]*string) *CreateDataProcessJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDataProcessJobTemplateResponse) SetStatusCode(v int32) *CreateDataProcessJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataProcessJobTemplateResponse) SetBody(v *CreateDataProcessJobTemplateResponseBody) *CreateDataProcessJobTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateDataProcessJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
