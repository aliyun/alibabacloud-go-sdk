// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchReplicationLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchReplicationLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchReplicationLinkResponse
	GetStatusCode() *int32
	SetBody(v *SwitchReplicationLinkResponseBody) *SwitchReplicationLinkResponse
	GetBody() *SwitchReplicationLinkResponseBody
}

type SwitchReplicationLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchReplicationLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchReplicationLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchReplicationLinkResponse) GoString() string {
	return s.String()
}

func (s *SwitchReplicationLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchReplicationLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchReplicationLinkResponse) GetBody() *SwitchReplicationLinkResponseBody {
	return s.Body
}

func (s *SwitchReplicationLinkResponse) SetHeaders(v map[string]*string) *SwitchReplicationLinkResponse {
	s.Headers = v
	return s
}

func (s *SwitchReplicationLinkResponse) SetStatusCode(v int32) *SwitchReplicationLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchReplicationLinkResponse) SetBody(v *SwitchReplicationLinkResponseBody) *SwitchReplicationLinkResponse {
	s.Body = v
	return s
}

func (s *SwitchReplicationLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
