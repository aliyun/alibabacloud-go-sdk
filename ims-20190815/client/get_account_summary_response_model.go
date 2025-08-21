// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountSummaryResponseBody) *GetAccountSummaryResponse
	GetBody() *GetAccountSummaryResponseBody
}

type GetAccountSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountSummaryResponse) GetBody() *GetAccountSummaryResponseBody {
	return s.Body
}

func (s *GetAccountSummaryResponse) SetHeaders(v map[string]*string) *GetAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSummaryResponse) SetStatusCode(v int32) *GetAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountSummaryResponse) SetBody(v *GetAccountSummaryResponseBody) *GetAccountSummaryResponse {
	s.Body = v
	return s
}

func (s *GetAccountSummaryResponse) Validate() error {
	return dara.Validate(s)
}
