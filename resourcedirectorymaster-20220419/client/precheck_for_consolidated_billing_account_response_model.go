// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckForConsolidatedBillingAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckForConsolidatedBillingAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckForConsolidatedBillingAccountResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckForConsolidatedBillingAccountResponseBody) *PrecheckForConsolidatedBillingAccountResponse
	GetBody() *PrecheckForConsolidatedBillingAccountResponseBody
}

type PrecheckForConsolidatedBillingAccountResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckForConsolidatedBillingAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponse) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckForConsolidatedBillingAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckForConsolidatedBillingAccountResponse) GetBody() *PrecheckForConsolidatedBillingAccountResponseBody {
	return s.Body
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetHeaders(v map[string]*string) *PrecheckForConsolidatedBillingAccountResponse {
	s.Headers = v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetStatusCode(v int32) *PrecheckForConsolidatedBillingAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetBody(v *PrecheckForConsolidatedBillingAccountResponseBody) *PrecheckForConsolidatedBillingAccountResponse {
	s.Body = v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponse) Validate() error {
	return dara.Validate(s)
}
