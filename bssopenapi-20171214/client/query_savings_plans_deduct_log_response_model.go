// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDeductLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySavingsPlansDeductLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySavingsPlansDeductLogResponse
	GetStatusCode() *int32
	SetBody(v *QuerySavingsPlansDeductLogResponseBody) *QuerySavingsPlansDeductLogResponse
	GetBody() *QuerySavingsPlansDeductLogResponseBody
}

type QuerySavingsPlansDeductLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySavingsPlansDeductLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySavingsPlansDeductLogResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponse) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySavingsPlansDeductLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySavingsPlansDeductLogResponse) GetBody() *QuerySavingsPlansDeductLogResponseBody {
	return s.Body
}

func (s *QuerySavingsPlansDeductLogResponse) SetHeaders(v map[string]*string) *QuerySavingsPlansDeductLogResponse {
	s.Headers = v
	return s
}

func (s *QuerySavingsPlansDeductLogResponse) SetStatusCode(v int32) *QuerySavingsPlansDeductLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponse) SetBody(v *QuerySavingsPlansDeductLogResponseBody) *QuerySavingsPlansDeductLogResponse {
	s.Body = v
	return s
}

func (s *QuerySavingsPlansDeductLogResponse) Validate() error {
	return dara.Validate(s)
}
