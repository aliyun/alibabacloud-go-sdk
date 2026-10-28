// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactGroupsResponseBody) *DeleteContactGroupsResponse
	GetBody() *DeleteContactGroupsResponseBody
}

type DeleteContactGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactGroupsResponse) GetBody() *DeleteContactGroupsResponseBody {
	return s.Body
}

func (s *DeleteContactGroupsResponse) SetHeaders(v map[string]*string) *DeleteContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactGroupsResponse) SetStatusCode(v int32) *DeleteContactGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactGroupsResponse) SetBody(v *DeleteContactGroupsResponseBody) *DeleteContactGroupsResponse {
	s.Body = v
	return s
}

func (s *DeleteContactGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
