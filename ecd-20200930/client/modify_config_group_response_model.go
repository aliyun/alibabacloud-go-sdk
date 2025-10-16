// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConfigGroupResponseBody) *ModifyConfigGroupResponse
	GetBody() *ModifyConfigGroupResponseBody
}

type ModifyConfigGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConfigGroupResponse) GetBody() *ModifyConfigGroupResponseBody {
	return s.Body
}

func (s *ModifyConfigGroupResponse) SetHeaders(v map[string]*string) *ModifyConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyConfigGroupResponse) SetStatusCode(v int32) *ModifyConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConfigGroupResponse) SetBody(v *ModifyConfigGroupResponseBody) *ModifyConfigGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyConfigGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
