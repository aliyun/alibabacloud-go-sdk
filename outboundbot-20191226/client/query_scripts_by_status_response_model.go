// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptsByStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryScriptsByStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryScriptsByStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryScriptsByStatusResponseBody) *QueryScriptsByStatusResponse
	GetBody() *QueryScriptsByStatusResponseBody
}

type QueryScriptsByStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScriptsByStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScriptsByStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptsByStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryScriptsByStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryScriptsByStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryScriptsByStatusResponse) GetBody() *QueryScriptsByStatusResponseBody {
	return s.Body
}

func (s *QueryScriptsByStatusResponse) SetHeaders(v map[string]*string) *QueryScriptsByStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryScriptsByStatusResponse) SetStatusCode(v int32) *QueryScriptsByStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScriptsByStatusResponse) SetBody(v *QueryScriptsByStatusResponseBody) *QueryScriptsByStatusResponse {
	s.Body = v
	return s
}

func (s *QueryScriptsByStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
