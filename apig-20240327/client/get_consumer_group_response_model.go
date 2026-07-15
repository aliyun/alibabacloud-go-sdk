// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerGroupResponseBody) *GetConsumerGroupResponse
	GetBody() *GetConsumerGroupResponseBody
}

type GetConsumerGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerGroupResponse) GetBody() *GetConsumerGroupResponseBody {
	return s.Body
}

func (s *GetConsumerGroupResponse) SetHeaders(v map[string]*string) *GetConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupResponse) SetStatusCode(v int32) *GetConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupResponse) SetBody(v *GetConsumerGroupResponseBody) *GetConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *GetConsumerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
