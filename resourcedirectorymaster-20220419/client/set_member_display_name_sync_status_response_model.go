// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDisplayNameSyncStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetMemberDisplayNameSyncStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetMemberDisplayNameSyncStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetMemberDisplayNameSyncStatusResponseBody) *SetMemberDisplayNameSyncStatusResponse
	GetBody() *SetMemberDisplayNameSyncStatusResponseBody
}

type SetMemberDisplayNameSyncStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMemberDisplayNameSyncStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMemberDisplayNameSyncStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDisplayNameSyncStatusResponse) GoString() string {
	return s.String()
}

func (s *SetMemberDisplayNameSyncStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetMemberDisplayNameSyncStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetMemberDisplayNameSyncStatusResponse) GetBody() *SetMemberDisplayNameSyncStatusResponseBody {
	return s.Body
}

func (s *SetMemberDisplayNameSyncStatusResponse) SetHeaders(v map[string]*string) *SetMemberDisplayNameSyncStatusResponse {
	s.Headers = v
	return s
}

func (s *SetMemberDisplayNameSyncStatusResponse) SetStatusCode(v int32) *SetMemberDisplayNameSyncStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMemberDisplayNameSyncStatusResponse) SetBody(v *SetMemberDisplayNameSyncStatusResponseBody) *SetMemberDisplayNameSyncStatusResponse {
	s.Body = v
	return s
}

func (s *SetMemberDisplayNameSyncStatusResponse) Validate() error {
	return dara.Validate(s)
}
