// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNoticeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNoticeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNoticeConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNoticeConfigResponseBody) *ModifyNoticeConfigResponse
	GetBody() *ModifyNoticeConfigResponseBody
}

type ModifyNoticeConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNoticeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNoticeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNoticeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNoticeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNoticeConfigResponse) GetBody() *ModifyNoticeConfigResponseBody {
	return s.Body
}

func (s *ModifyNoticeConfigResponse) SetHeaders(v map[string]*string) *ModifyNoticeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNoticeConfigResponse) SetStatusCode(v int32) *ModifyNoticeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNoticeConfigResponse) SetBody(v *ModifyNoticeConfigResponseBody) *ModifyNoticeConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyNoticeConfigResponse) Validate() error {
	return dara.Validate(s)
}
