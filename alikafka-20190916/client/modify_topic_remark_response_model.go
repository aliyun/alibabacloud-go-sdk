// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTopicRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTopicRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTopicRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTopicRemarkResponseBody) *ModifyTopicRemarkResponse
	GetBody() *ModifyTopicRemarkResponseBody
}

type ModifyTopicRemarkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTopicRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTopicRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTopicRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTopicRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTopicRemarkResponse) GetBody() *ModifyTopicRemarkResponseBody {
	return s.Body
}

func (s *ModifyTopicRemarkResponse) SetHeaders(v map[string]*string) *ModifyTopicRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyTopicRemarkResponse) SetStatusCode(v int32) *ModifyTopicRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTopicRemarkResponse) SetBody(v *ModifyTopicRemarkResponseBody) *ModifyTopicRemarkResponse {
	s.Body = v
	return s
}

func (s *ModifyTopicRemarkResponse) Validate() error {
	return dara.Validate(s)
}
