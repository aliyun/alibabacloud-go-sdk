// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConsumerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConsumerResponseBody) *ModifyConsumerResponse
	GetBody() *ModifyConsumerResponseBody
}

type ModifyConsumerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConsumerResponse) GetBody() *ModifyConsumerResponseBody {
	return s.Body
}

func (s *ModifyConsumerResponse) SetHeaders(v map[string]*string) *ModifyConsumerResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerResponse) SetStatusCode(v int32) *ModifyConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumerResponse) SetBody(v *ModifyConsumerResponseBody) *ModifyConsumerResponse {
	s.Body = v
	return s
}

func (s *ModifyConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
