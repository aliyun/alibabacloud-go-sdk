// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntelligentStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntelligentStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntelligentStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntelligentStrategyResponseBody) *CreateIntelligentStrategyResponse
	GetBody() *CreateIntelligentStrategyResponseBody
}

type CreateIntelligentStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntelligentStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntelligentStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntelligentStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateIntelligentStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntelligentStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntelligentStrategyResponse) GetBody() *CreateIntelligentStrategyResponseBody {
	return s.Body
}

func (s *CreateIntelligentStrategyResponse) SetHeaders(v map[string]*string) *CreateIntelligentStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateIntelligentStrategyResponse) SetStatusCode(v int32) *CreateIntelligentStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntelligentStrategyResponse) SetBody(v *CreateIntelligentStrategyResponseBody) *CreateIntelligentStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateIntelligentStrategyResponse) Validate() error {
	return dara.Validate(s)
}
