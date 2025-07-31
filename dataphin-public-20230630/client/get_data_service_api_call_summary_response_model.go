// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiCallSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiCallSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiCallSummaryResponseBody) *GetDataServiceApiCallSummaryResponse
	GetBody() *GetDataServiceApiCallSummaryResponseBody
}

type GetDataServiceApiCallSummaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiCallSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiCallSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiCallSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiCallSummaryResponse) GetBody() *GetDataServiceApiCallSummaryResponseBody {
	return s.Body
}

func (s *GetDataServiceApiCallSummaryResponse) SetHeaders(v map[string]*string) *GetDataServiceApiCallSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiCallSummaryResponse) SetStatusCode(v int32) *GetDataServiceApiCallSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponse) SetBody(v *GetDataServiceApiCallSummaryResponseBody) *GetDataServiceApiCallSummaryResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiCallSummaryResponse) Validate() error {
	return dara.Validate(s)
}
