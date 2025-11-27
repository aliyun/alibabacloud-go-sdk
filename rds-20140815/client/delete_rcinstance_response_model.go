// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCInstanceResponseBody) *DeleteRCInstanceResponse
	GetBody() *DeleteRCInstanceResponseBody
}

type DeleteRCInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCInstanceResponse) GetBody() *DeleteRCInstanceResponseBody {
	return s.Body
}

func (s *DeleteRCInstanceResponse) SetHeaders(v map[string]*string) *DeleteRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCInstanceResponse) SetStatusCode(v int32) *DeleteRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCInstanceResponse) SetBody(v *DeleteRCInstanceResponseBody) *DeleteRCInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
