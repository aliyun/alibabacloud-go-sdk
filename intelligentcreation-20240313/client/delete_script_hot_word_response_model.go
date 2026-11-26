// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptHotWordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScriptHotWordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScriptHotWordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScriptHotWordResponseBody) *DeleteScriptHotWordResponse
	GetBody() *DeleteScriptHotWordResponseBody
}

type DeleteScriptHotWordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScriptHotWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScriptHotWordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptHotWordResponse) GoString() string {
	return s.String()
}

func (s *DeleteScriptHotWordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScriptHotWordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScriptHotWordResponse) GetBody() *DeleteScriptHotWordResponseBody {
	return s.Body
}

func (s *DeleteScriptHotWordResponse) SetHeaders(v map[string]*string) *DeleteScriptHotWordResponse {
	s.Headers = v
	return s
}

func (s *DeleteScriptHotWordResponse) SetStatusCode(v int32) *DeleteScriptHotWordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScriptHotWordResponse) SetBody(v *DeleteScriptHotWordResponseBody) *DeleteScriptHotWordResponse {
	s.Body = v
	return s
}

func (s *DeleteScriptHotWordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
