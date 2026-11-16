// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumeSummaryResponseBody) *ListConsumeSummaryResponse
	GetBody() *ListConsumeSummaryResponseBody
}

type ListConsumeSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumeSummaryResponse) GetBody() *ListConsumeSummaryResponseBody {
	return s.Body
}

func (s *ListConsumeSummaryResponse) SetHeaders(v map[string]*string) *ListConsumeSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListConsumeSummaryResponse) SetStatusCode(v int32) *ListConsumeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumeSummaryResponse) SetBody(v *ListConsumeSummaryResponseBody) *ListConsumeSummaryResponse {
	s.Body = v
	return s
}

func (s *ListConsumeSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
