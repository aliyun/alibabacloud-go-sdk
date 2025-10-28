// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterMemberResponseBody) *DeleteClusterMemberResponse
	GetBody() *DeleteClusterMemberResponseBody
}

type DeleteClusterMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterMemberResponse) GetBody() *DeleteClusterMemberResponseBody {
	return s.Body
}

func (s *DeleteClusterMemberResponse) SetHeaders(v map[string]*string) *DeleteClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterMemberResponse) SetStatusCode(v int32) *DeleteClusterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterMemberResponse) SetBody(v *DeleteClusterMemberResponseBody) *DeleteClusterMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteClusterMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
