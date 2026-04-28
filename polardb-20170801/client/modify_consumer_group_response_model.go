// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConsumerGroupResponseBody) *ModifyConsumerGroupResponse
	GetBody() *ModifyConsumerGroupResponseBody
}

type ModifyConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConsumerGroupResponse) GetBody() *ModifyConsumerGroupResponseBody {
	return s.Body
}

func (s *ModifyConsumerGroupResponse) SetHeaders(v map[string]*string) *ModifyConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerGroupResponse) SetStatusCode(v int32) *ModifyConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumerGroupResponse) SetBody(v *ModifyConsumerGroupResponseBody) *ModifyConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyConsumerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
