// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAutomateResponseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostAutomateResponseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostAutomateResponseConfigResponse
	GetStatusCode() *int32
	SetBody(v *PostAutomateResponseConfigResponseBody) *PostAutomateResponseConfigResponse
	GetBody() *PostAutomateResponseConfigResponseBody
}

type PostAutomateResponseConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostAutomateResponseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PostAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostAutomateResponseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostAutomateResponseConfigResponse) GetBody() *PostAutomateResponseConfigResponseBody {
	return s.Body
}

func (s *PostAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *PostAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetStatusCode(v int32) *PostAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetBody(v *PostAutomateResponseConfigResponseBody) *PostAutomateResponseConfigResponse {
	s.Body = v
	return s
}

func (s *PostAutomateResponseConfigResponse) Validate() error {
	return dara.Validate(s)
}
