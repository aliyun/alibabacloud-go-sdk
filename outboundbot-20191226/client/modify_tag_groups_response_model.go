// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTagGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTagGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTagGroupsResponseBody) *ModifyTagGroupsResponse
	GetBody() *ModifyTagGroupsResponseBody
}

type ModifyTagGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTagGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTagGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyTagGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTagGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTagGroupsResponse) GetBody() *ModifyTagGroupsResponseBody {
	return s.Body
}

func (s *ModifyTagGroupsResponse) SetHeaders(v map[string]*string) *ModifyTagGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyTagGroupsResponse) SetStatusCode(v int32) *ModifyTagGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTagGroupsResponse) SetBody(v *ModifyTagGroupsResponseBody) *ModifyTagGroupsResponse {
	s.Body = v
	return s
}

func (s *ModifyTagGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
