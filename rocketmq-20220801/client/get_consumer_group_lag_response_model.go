// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupLagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerGroupLagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerGroupLagResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerGroupLagResponseBody) *GetConsumerGroupLagResponse
	GetBody() *GetConsumerGroupLagResponseBody
}

type GetConsumerGroupLagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupLagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupLagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupLagResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerGroupLagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerGroupLagResponse) GetBody() *GetConsumerGroupLagResponseBody {
	return s.Body
}

func (s *GetConsumerGroupLagResponse) SetHeaders(v map[string]*string) *GetConsumerGroupLagResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupLagResponse) SetStatusCode(v int32) *GetConsumerGroupLagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupLagResponse) SetBody(v *GetConsumerGroupLagResponseBody) *GetConsumerGroupLagResponse {
	s.Body = v
	return s
}

func (s *GetConsumerGroupLagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
