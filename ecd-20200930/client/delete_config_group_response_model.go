// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigGroupResponseBody) *DeleteConfigGroupResponse
	GetBody() *DeleteConfigGroupResponseBody
}

type DeleteConfigGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigGroupResponse) GetBody() *DeleteConfigGroupResponseBody {
	return s.Body
}

func (s *DeleteConfigGroupResponse) SetHeaders(v map[string]*string) *DeleteConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigGroupResponse) SetStatusCode(v int32) *DeleteConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigGroupResponse) SetBody(v *DeleteConfigGroupResponseBody) *DeleteConfigGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
