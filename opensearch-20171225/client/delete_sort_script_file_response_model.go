// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSortScriptFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSortScriptFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSortScriptFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSortScriptFileResponseBody) *DeleteSortScriptFileResponse
	GetBody() *DeleteSortScriptFileResponseBody
}

type DeleteSortScriptFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSortScriptFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSortScriptFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSortScriptFileResponse) GetBody() *DeleteSortScriptFileResponseBody {
	return s.Body
}

func (s *DeleteSortScriptFileResponse) SetHeaders(v map[string]*string) *DeleteSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptFileResponse) SetStatusCode(v int32) *DeleteSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSortScriptFileResponse) SetBody(v *DeleteSortScriptFileResponseBody) *DeleteSortScriptFileResponse {
	s.Body = v
	return s
}

func (s *DeleteSortScriptFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
