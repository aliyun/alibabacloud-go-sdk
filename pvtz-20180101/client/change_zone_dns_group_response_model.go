// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeZoneDnsGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeZoneDnsGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeZoneDnsGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeZoneDnsGroupResponseBody) *ChangeZoneDnsGroupResponse
	GetBody() *ChangeZoneDnsGroupResponseBody
}

type ChangeZoneDnsGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeZoneDnsGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeZoneDnsGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeZoneDnsGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeZoneDnsGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeZoneDnsGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeZoneDnsGroupResponse) GetBody() *ChangeZoneDnsGroupResponseBody {
	return s.Body
}

func (s *ChangeZoneDnsGroupResponse) SetHeaders(v map[string]*string) *ChangeZoneDnsGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeZoneDnsGroupResponse) SetStatusCode(v int32) *ChangeZoneDnsGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeZoneDnsGroupResponse) SetBody(v *ChangeZoneDnsGroupResponseBody) *ChangeZoneDnsGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeZoneDnsGroupResponse) Validate() error {
	return dara.Validate(s)
}
