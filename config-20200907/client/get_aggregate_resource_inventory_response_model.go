// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceInventoryResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceInventoryResponseBody) *GetAggregateResourceInventoryResponse
	GetBody() *GetAggregateResourceInventoryResponseBody
}

type GetAggregateResourceInventoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceInventoryResponse) GetBody() *GetAggregateResourceInventoryResponseBody {
	return s.Body
}

func (s *GetAggregateResourceInventoryResponse) SetHeaders(v map[string]*string) *GetAggregateResourceInventoryResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceInventoryResponse) SetStatusCode(v int32) *GetAggregateResourceInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceInventoryResponse) SetBody(v *GetAggregateResourceInventoryResponseBody) *GetAggregateResourceInventoryResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
