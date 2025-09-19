// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerlessAuthSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServerlessAuthSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServerlessAuthSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetServerlessAuthSummaryResponseBody) *GetServerlessAuthSummaryResponse
	GetBody() *GetServerlessAuthSummaryResponseBody
}

type GetServerlessAuthSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServerlessAuthSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServerlessAuthSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServerlessAuthSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetServerlessAuthSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServerlessAuthSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServerlessAuthSummaryResponse) GetBody() *GetServerlessAuthSummaryResponseBody {
	return s.Body
}

func (s *GetServerlessAuthSummaryResponse) SetHeaders(v map[string]*string) *GetServerlessAuthSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetServerlessAuthSummaryResponse) SetStatusCode(v int32) *GetServerlessAuthSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServerlessAuthSummaryResponse) SetBody(v *GetServerlessAuthSummaryResponseBody) *GetServerlessAuthSummaryResponse {
	s.Body = v
	return s
}

func (s *GetServerlessAuthSummaryResponse) Validate() error {
	return dara.Validate(s)
}
