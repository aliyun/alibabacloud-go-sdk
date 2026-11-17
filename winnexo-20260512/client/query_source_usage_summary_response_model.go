// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySourceUsageSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySourceUsageSummaryResponse
	GetStatusCode() *int32
	SetBody(v *QuerySourceUsageSummaryResponseBody) *QuerySourceUsageSummaryResponse
	GetBody() *QuerySourceUsageSummaryResponseBody
}

type QuerySourceUsageSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySourceUsageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySourceUsageSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageSummaryResponse) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySourceUsageSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySourceUsageSummaryResponse) GetBody() *QuerySourceUsageSummaryResponseBody {
	return s.Body
}

func (s *QuerySourceUsageSummaryResponse) SetHeaders(v map[string]*string) *QuerySourceUsageSummaryResponse {
	s.Headers = v
	return s
}

func (s *QuerySourceUsageSummaryResponse) SetStatusCode(v int32) *QuerySourceUsageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySourceUsageSummaryResponse) SetBody(v *QuerySourceUsageSummaryResponseBody) *QuerySourceUsageSummaryResponse {
	s.Body = v
	return s
}

func (s *QuerySourceUsageSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
