// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodPackagingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodPackagingGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetVodPackagingGroupResponseBody) *GetVodPackagingGroupResponse
	GetBody() *GetVodPackagingGroupResponseBody
}

type GetVodPackagingGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVodPackagingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodPackagingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingGroupResponse) GoString() string {
	return s.String()
}

func (s *GetVodPackagingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodPackagingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodPackagingGroupResponse) GetBody() *GetVodPackagingGroupResponseBody {
	return s.Body
}

func (s *GetVodPackagingGroupResponse) SetHeaders(v map[string]*string) *GetVodPackagingGroupResponse {
	s.Headers = v
	return s
}

func (s *GetVodPackagingGroupResponse) SetStatusCode(v int32) *GetVodPackagingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodPackagingGroupResponse) SetBody(v *GetVodPackagingGroupResponseBody) *GetVodPackagingGroupResponse {
	s.Body = v
	return s
}

func (s *GetVodPackagingGroupResponse) Validate() error {
	return dara.Validate(s)
}
