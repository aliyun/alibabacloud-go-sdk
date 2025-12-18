// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateResourceInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAggregateResourceInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAggregateResourceInventoryResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAggregateResourceInventoryResponseBody) *GenerateAggregateResourceInventoryResponse
	GetBody() *GenerateAggregateResourceInventoryResponseBody
}

type GenerateAggregateResourceInventoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAggregateResourceInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAggregateResourceInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateResourceInventoryResponse) GoString() string {
	return s.String()
}

func (s *GenerateAggregateResourceInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAggregateResourceInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAggregateResourceInventoryResponse) GetBody() *GenerateAggregateResourceInventoryResponseBody {
	return s.Body
}

func (s *GenerateAggregateResourceInventoryResponse) SetHeaders(v map[string]*string) *GenerateAggregateResourceInventoryResponse {
	s.Headers = v
	return s
}

func (s *GenerateAggregateResourceInventoryResponse) SetStatusCode(v int32) *GenerateAggregateResourceInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAggregateResourceInventoryResponse) SetBody(v *GenerateAggregateResourceInventoryResponseBody) *GenerateAggregateResourceInventoryResponse {
	s.Body = v
	return s
}

func (s *GenerateAggregateResourceInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
