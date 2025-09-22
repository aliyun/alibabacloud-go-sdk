// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventorySummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGwpInventorySummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGwpInventorySummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetGwpInventorySummaryResponseBody) *GetGwpInventorySummaryResponse
	GetBody() *GetGwpInventorySummaryResponseBody
}

type GetGwpInventorySummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpInventorySummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpInventorySummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventorySummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGwpInventorySummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGwpInventorySummaryResponse) GetBody() *GetGwpInventorySummaryResponseBody {
	return s.Body
}

func (s *GetGwpInventorySummaryResponse) SetHeaders(v map[string]*string) *GetGwpInventorySummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGwpInventorySummaryResponse) SetStatusCode(v int32) *GetGwpInventorySummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpInventorySummaryResponse) SetBody(v *GetGwpInventorySummaryResponseBody) *GetGwpInventorySummaryResponse {
	s.Body = v
	return s
}

func (s *GetGwpInventorySummaryResponse) Validate() error {
	return dara.Validate(s)
}
