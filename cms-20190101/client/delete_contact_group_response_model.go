// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactGroupResponseBody) *DeleteContactGroupResponse
	GetBody() *DeleteContactGroupResponseBody
}

type DeleteContactGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactGroupResponse) GetBody() *DeleteContactGroupResponseBody {
	return s.Body
}

func (s *DeleteContactGroupResponse) SetHeaders(v map[string]*string) *DeleteContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactGroupResponse) SetStatusCode(v int32) *DeleteContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactGroupResponse) SetBody(v *DeleteContactGroupResponseBody) *DeleteContactGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
