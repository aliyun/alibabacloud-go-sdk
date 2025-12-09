// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootRenderingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootRenderingServerResponse
	GetStatusCode() *int32
	SetBody(v *RebootRenderingServerResponseBody) *RebootRenderingServerResponse
	GetBody() *RebootRenderingServerResponseBody
}

type RebootRenderingServerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootRenderingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootRenderingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerResponse) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootRenderingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootRenderingServerResponse) GetBody() *RebootRenderingServerResponseBody {
	return s.Body
}

func (s *RebootRenderingServerResponse) SetHeaders(v map[string]*string) *RebootRenderingServerResponse {
	s.Headers = v
	return s
}

func (s *RebootRenderingServerResponse) SetStatusCode(v int32) *RebootRenderingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootRenderingServerResponse) SetBody(v *RebootRenderingServerResponseBody) *RebootRenderingServerResponse {
	s.Body = v
	return s
}

func (s *RebootRenderingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
