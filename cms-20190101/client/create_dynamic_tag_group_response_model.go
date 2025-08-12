// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDynamicTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDynamicTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDynamicTagGroupResponseBody) *CreateDynamicTagGroupResponse
	GetBody() *CreateDynamicTagGroupResponseBody
}

type CreateDynamicTagGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDynamicTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDynamicTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicTagGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDynamicTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDynamicTagGroupResponse) GetBody() *CreateDynamicTagGroupResponseBody {
	return s.Body
}

func (s *CreateDynamicTagGroupResponse) SetHeaders(v map[string]*string) *CreateDynamicTagGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDynamicTagGroupResponse) SetStatusCode(v int32) *CreateDynamicTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDynamicTagGroupResponse) SetBody(v *CreateDynamicTagGroupResponseBody) *CreateDynamicTagGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDynamicTagGroupResponse) Validate() error {
	return dara.Validate(s)
}
