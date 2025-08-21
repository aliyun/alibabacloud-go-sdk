// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalAttackMaxFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTotalAttackMaxFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTotalAttackMaxFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTotalAttackMaxFlowResponseBody) *DescribeTotalAttackMaxFlowResponse
	GetBody() *DescribeTotalAttackMaxFlowResponseBody
}

type DescribeTotalAttackMaxFlowResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTotalAttackMaxFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTotalAttackMaxFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalAttackMaxFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeTotalAttackMaxFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTotalAttackMaxFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTotalAttackMaxFlowResponse) GetBody() *DescribeTotalAttackMaxFlowResponseBody {
	return s.Body
}

func (s *DescribeTotalAttackMaxFlowResponse) SetHeaders(v map[string]*string) *DescribeTotalAttackMaxFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponse) SetStatusCode(v int32) *DescribeTotalAttackMaxFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponse) SetBody(v *DescribeTotalAttackMaxFlowResponseBody) *DescribeTotalAttackMaxFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponse) Validate() error {
	return dara.Validate(s)
}
