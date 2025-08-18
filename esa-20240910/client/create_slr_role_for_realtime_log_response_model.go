// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrRoleForRealtimeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlrRoleForRealtimeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlrRoleForRealtimeLogResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlrRoleForRealtimeLogResponseBody) *CreateSlrRoleForRealtimeLogResponse
	GetBody() *CreateSlrRoleForRealtimeLogResponseBody
}

type CreateSlrRoleForRealtimeLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlrRoleForRealtimeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlrRoleForRealtimeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrRoleForRealtimeLogResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleForRealtimeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlrRoleForRealtimeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlrRoleForRealtimeLogResponse) GetBody() *CreateSlrRoleForRealtimeLogResponseBody {
	return s.Body
}

func (s *CreateSlrRoleForRealtimeLogResponse) SetHeaders(v map[string]*string) *CreateSlrRoleForRealtimeLogResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrRoleForRealtimeLogResponse) SetStatusCode(v int32) *CreateSlrRoleForRealtimeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlrRoleForRealtimeLogResponse) SetBody(v *CreateSlrRoleForRealtimeLogResponseBody) *CreateSlrRoleForRealtimeLogResponse {
	s.Body = v
	return s
}

func (s *CreateSlrRoleForRealtimeLogResponse) Validate() error {
	return dara.Validate(s)
}
