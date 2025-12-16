// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *RenewAppGroupResponseBody) *RenewAppGroupResponse
	GetBody() *RenewAppGroupResponseBody
}

type RenewAppGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAppGroupResponse) GetBody() *RenewAppGroupResponseBody {
	return s.Body
}

func (s *RenewAppGroupResponse) SetHeaders(v map[string]*string) *RenewAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppGroupResponse) SetStatusCode(v int32) *RenewAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppGroupResponse) SetBody(v *RenewAppGroupResponseBody) *RenewAppGroupResponse {
	s.Body = v
	return s
}

func (s *RenewAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
