// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckSummaryResponseBody) *GetCheckSummaryResponse
	GetBody() *GetCheckSummaryResponseBody
}

type GetCheckSummaryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckSummaryResponse) GetBody() *GetCheckSummaryResponseBody {
	return s.Body
}

func (s *GetCheckSummaryResponse) SetHeaders(v map[string]*string) *GetCheckSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetCheckSummaryResponse) SetStatusCode(v int32) *GetCheckSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckSummaryResponse) SetBody(v *GetCheckSummaryResponseBody) *GetCheckSummaryResponse {
	s.Body = v
	return s
}

func (s *GetCheckSummaryResponse) Validate() error {
	return dara.Validate(s)
}
