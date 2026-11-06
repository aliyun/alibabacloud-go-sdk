// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataProcessJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataProcessJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataProcessJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataProcessJobTemplateResponseBody) *DeleteDataProcessJobTemplateResponse
	GetBody() *DeleteDataProcessJobTemplateResponseBody
}

type DeleteDataProcessJobTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataProcessJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataProcessJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataProcessJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataProcessJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataProcessJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataProcessJobTemplateResponse) GetBody() *DeleteDataProcessJobTemplateResponseBody {
	return s.Body
}

func (s *DeleteDataProcessJobTemplateResponse) SetHeaders(v map[string]*string) *DeleteDataProcessJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataProcessJobTemplateResponse) SetStatusCode(v int32) *DeleteDataProcessJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataProcessJobTemplateResponse) SetBody(v *DeleteDataProcessJobTemplateResponseBody) *DeleteDataProcessJobTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDataProcessJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
