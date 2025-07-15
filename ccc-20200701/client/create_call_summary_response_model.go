// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallSummaryResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallSummaryResponseBody) *CreateCallSummaryResponse
	GetBody() *CreateCallSummaryResponseBody
}

type CreateCallSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallSummaryResponse) GoString() string {
	return s.String()
}

func (s *CreateCallSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallSummaryResponse) GetBody() *CreateCallSummaryResponseBody {
	return s.Body
}

func (s *CreateCallSummaryResponse) SetHeaders(v map[string]*string) *CreateCallSummaryResponse {
	s.Headers = v
	return s
}

func (s *CreateCallSummaryResponse) SetStatusCode(v int32) *CreateCallSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallSummaryResponse) SetBody(v *CreateCallSummaryResponseBody) *CreateCallSummaryResponse {
	s.Body = v
	return s
}

func (s *CreateCallSummaryResponse) Validate() error {
	return dara.Validate(s)
}
