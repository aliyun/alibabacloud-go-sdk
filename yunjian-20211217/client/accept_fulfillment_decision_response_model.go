// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptFulfillmentDecisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptFulfillmentDecisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptFulfillmentDecisionResponse
	GetStatusCode() *int32
	SetBody(v *AcceptFulfillmentDecisionResponseBody) *AcceptFulfillmentDecisionResponse
	GetBody() *AcceptFulfillmentDecisionResponseBody
}

type AcceptFulfillmentDecisionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptFulfillmentDecisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptFulfillmentDecisionResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptFulfillmentDecisionResponse) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptFulfillmentDecisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptFulfillmentDecisionResponse) GetBody() *AcceptFulfillmentDecisionResponseBody {
	return s.Body
}

func (s *AcceptFulfillmentDecisionResponse) SetHeaders(v map[string]*string) *AcceptFulfillmentDecisionResponse {
	s.Headers = v
	return s
}

func (s *AcceptFulfillmentDecisionResponse) SetStatusCode(v int32) *AcceptFulfillmentDecisionResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponse) SetBody(v *AcceptFulfillmentDecisionResponseBody) *AcceptFulfillmentDecisionResponse {
	s.Body = v
	return s
}

func (s *AcceptFulfillmentDecisionResponse) Validate() error {
	return dara.Validate(s)
}
