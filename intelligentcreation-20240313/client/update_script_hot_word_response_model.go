// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptHotWordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScriptHotWordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScriptHotWordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScriptHotWordResponseBody) *UpdateScriptHotWordResponse
	GetBody() *UpdateScriptHotWordResponseBody
}

type UpdateScriptHotWordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScriptHotWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScriptHotWordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptHotWordResponse) GoString() string {
	return s.String()
}

func (s *UpdateScriptHotWordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScriptHotWordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScriptHotWordResponse) GetBody() *UpdateScriptHotWordResponseBody {
	return s.Body
}

func (s *UpdateScriptHotWordResponse) SetHeaders(v map[string]*string) *UpdateScriptHotWordResponse {
	s.Headers = v
	return s
}

func (s *UpdateScriptHotWordResponse) SetStatusCode(v int32) *UpdateScriptHotWordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScriptHotWordResponse) SetBody(v *UpdateScriptHotWordResponseBody) *UpdateScriptHotWordResponse {
	s.Body = v
	return s
}

func (s *UpdateScriptHotWordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
