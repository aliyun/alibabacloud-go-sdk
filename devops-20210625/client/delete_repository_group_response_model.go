// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepositoryGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepositoryGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepositoryGroupResponseBody) *DeleteRepositoryGroupResponse
	GetBody() *DeleteRepositoryGroupResponseBody
}

type DeleteRepositoryGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepositoryGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepositoryGroupResponse) GetBody() *DeleteRepositoryGroupResponseBody {
	return s.Body
}

func (s *DeleteRepositoryGroupResponse) SetHeaders(v map[string]*string) *DeleteRepositoryGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryGroupResponse) SetStatusCode(v int32) *DeleteRepositoryGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryGroupResponse) SetBody(v *DeleteRepositoryGroupResponseBody) *DeleteRepositoryGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteRepositoryGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
