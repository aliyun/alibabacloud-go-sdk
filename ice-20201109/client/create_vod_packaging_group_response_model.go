// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVodPackagingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVodPackagingGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateVodPackagingGroupResponseBody) *CreateVodPackagingGroupResponse
	GetBody() *CreateVodPackagingGroupResponseBody
}

type CreateVodPackagingGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVodPackagingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVodPackagingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVodPackagingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVodPackagingGroupResponse) GetBody() *CreateVodPackagingGroupResponseBody {
	return s.Body
}

func (s *CreateVodPackagingGroupResponse) SetHeaders(v map[string]*string) *CreateVodPackagingGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVodPackagingGroupResponse) SetStatusCode(v int32) *CreateVodPackagingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVodPackagingGroupResponse) SetBody(v *CreateVodPackagingGroupResponseBody) *CreateVodPackagingGroupResponse {
	s.Body = v
	return s
}

func (s *CreateVodPackagingGroupResponse) Validate() error {
	return dara.Validate(s)
}
