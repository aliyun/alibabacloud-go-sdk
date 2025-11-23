// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardGroupResponseBody) *DeleteStandardGroupResponse
	GetBody() *DeleteStandardGroupResponseBody
}

type DeleteStandardGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardGroupResponse) GetBody() *DeleteStandardGroupResponseBody {
	return s.Body
}

func (s *DeleteStandardGroupResponse) SetHeaders(v map[string]*string) *DeleteStandardGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardGroupResponse) SetStatusCode(v int32) *DeleteStandardGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardGroupResponse) SetBody(v *DeleteStandardGroupResponseBody) *DeleteStandardGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
