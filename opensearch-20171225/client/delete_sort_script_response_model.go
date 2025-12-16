// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSortScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSortScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSortScriptResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSortScriptResponseBody) *DeleteSortScriptResponse
	GetBody() *DeleteSortScriptResponseBody
}

type DeleteSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSortScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSortScriptResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSortScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSortScriptResponse) GetBody() *DeleteSortScriptResponseBody {
	return s.Body
}

func (s *DeleteSortScriptResponse) SetHeaders(v map[string]*string) *DeleteSortScriptResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptResponse) SetStatusCode(v int32) *DeleteSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSortScriptResponse) SetBody(v *DeleteSortScriptResponseBody) *DeleteSortScriptResponse {
	s.Body = v
	return s
}

func (s *DeleteSortScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
