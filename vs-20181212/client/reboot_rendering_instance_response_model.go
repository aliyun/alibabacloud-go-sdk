// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebootRenderingInstanceResponseBody) *RebootRenderingInstanceResponse
	GetBody() *RebootRenderingInstanceResponseBody
}

type RebootRenderingInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootRenderingInstanceResponse) GetBody() *RebootRenderingInstanceResponseBody {
	return s.Body
}

func (s *RebootRenderingInstanceResponse) SetHeaders(v map[string]*string) *RebootRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootRenderingInstanceResponse) SetStatusCode(v int32) *RebootRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootRenderingInstanceResponse) SetBody(v *RebootRenderingInstanceResponseBody) *RebootRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *RebootRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
