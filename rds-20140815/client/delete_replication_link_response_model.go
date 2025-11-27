// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReplicationLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReplicationLinkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReplicationLinkResponseBody) *DeleteReplicationLinkResponse
	GetBody() *DeleteReplicationLinkResponseBody
}

type DeleteReplicationLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReplicationLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReplicationLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteReplicationLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReplicationLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReplicationLinkResponse) GetBody() *DeleteReplicationLinkResponseBody {
	return s.Body
}

func (s *DeleteReplicationLinkResponse) SetHeaders(v map[string]*string) *DeleteReplicationLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteReplicationLinkResponse) SetStatusCode(v int32) *DeleteReplicationLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReplicationLinkResponse) SetBody(v *DeleteReplicationLinkResponseBody) *DeleteReplicationLinkResponse {
	s.Body = v
	return s
}

func (s *DeleteReplicationLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
