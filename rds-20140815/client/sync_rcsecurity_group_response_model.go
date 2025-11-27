// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncRCSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncRCSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *SyncRCSecurityGroupResponseBody) *SyncRCSecurityGroupResponse
	GetBody() *SyncRCSecurityGroupResponseBody
}

type SyncRCSecurityGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncRCSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncRCSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncRCSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *SyncRCSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncRCSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncRCSecurityGroupResponse) GetBody() *SyncRCSecurityGroupResponseBody {
	return s.Body
}

func (s *SyncRCSecurityGroupResponse) SetHeaders(v map[string]*string) *SyncRCSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *SyncRCSecurityGroupResponse) SetStatusCode(v int32) *SyncRCSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncRCSecurityGroupResponse) SetBody(v *SyncRCSecurityGroupResponseBody) *SyncRCSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *SyncRCSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
