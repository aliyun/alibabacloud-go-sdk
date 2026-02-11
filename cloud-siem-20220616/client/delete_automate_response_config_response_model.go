// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutomateResponseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutomateResponseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutomateResponseConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutomateResponseConfigResponseBody) *DeleteAutomateResponseConfigResponse
	GetBody() *DeleteAutomateResponseConfigResponseBody
}

type DeleteAutomateResponseConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutomateResponseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutomateResponseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutomateResponseConfigResponse) GetBody() *DeleteAutomateResponseConfigResponseBody {
	return s.Body
}

func (s *DeleteAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *DeleteAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetStatusCode(v int32) *DeleteAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetBody(v *DeleteAutomateResponseConfigResponseBody) *DeleteAutomateResponseConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
