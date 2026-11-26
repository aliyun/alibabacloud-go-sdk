// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScriptHotWordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddScriptHotWordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddScriptHotWordResponse
	GetStatusCode() *int32
	SetBody(v *AddScriptHotWordResponseBody) *AddScriptHotWordResponse
	GetBody() *AddScriptHotWordResponseBody
}

type AddScriptHotWordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddScriptHotWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddScriptHotWordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddScriptHotWordResponse) GoString() string {
	return s.String()
}

func (s *AddScriptHotWordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddScriptHotWordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddScriptHotWordResponse) GetBody() *AddScriptHotWordResponseBody {
	return s.Body
}

func (s *AddScriptHotWordResponse) SetHeaders(v map[string]*string) *AddScriptHotWordResponse {
	s.Headers = v
	return s
}

func (s *AddScriptHotWordResponse) SetStatusCode(v int32) *AddScriptHotWordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddScriptHotWordResponse) SetBody(v *AddScriptHotWordResponseBody) *AddScriptHotWordResponse {
	s.Body = v
	return s
}

func (s *AddScriptHotWordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
