// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildReplicationLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildReplicationLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildReplicationLinkResponse
	GetStatusCode() *int32
	SetBody(v *RebuildReplicationLinkResponseBody) *RebuildReplicationLinkResponse
	GetBody() *RebuildReplicationLinkResponseBody
}

type RebuildReplicationLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildReplicationLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildReplicationLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildReplicationLinkResponse) GoString() string {
	return s.String()
}

func (s *RebuildReplicationLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildReplicationLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildReplicationLinkResponse) GetBody() *RebuildReplicationLinkResponseBody {
	return s.Body
}

func (s *RebuildReplicationLinkResponse) SetHeaders(v map[string]*string) *RebuildReplicationLinkResponse {
	s.Headers = v
	return s
}

func (s *RebuildReplicationLinkResponse) SetStatusCode(v int32) *RebuildReplicationLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildReplicationLinkResponse) SetBody(v *RebuildReplicationLinkResponseBody) *RebuildReplicationLinkResponse {
	s.Body = v
	return s
}

func (s *RebuildReplicationLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
