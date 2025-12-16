// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppGroupResponseBody) *ModifyAppGroupResponse
	GetBody() *ModifyAppGroupResponseBody
}

type ModifyAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppGroupResponse) GetBody() *ModifyAppGroupResponseBody {
	return s.Body
}

func (s *ModifyAppGroupResponse) SetHeaders(v map[string]*string) *ModifyAppGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupResponse) SetStatusCode(v int32) *ModifyAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppGroupResponse) SetBody(v *ModifyAppGroupResponseBody) *ModifyAppGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
