// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkTaskSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkTaskSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkTaskSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkTaskSummaryResponseBody) *GetWorkTaskSummaryResponse
	GetBody() *GetWorkTaskSummaryResponseBody
}

type GetWorkTaskSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkTaskSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkTaskSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkTaskSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkTaskSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkTaskSummaryResponse) GetBody() *GetWorkTaskSummaryResponseBody {
	return s.Body
}

func (s *GetWorkTaskSummaryResponse) SetHeaders(v map[string]*string) *GetWorkTaskSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetStatusCode(v int32) *GetWorkTaskSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetBody(v *GetWorkTaskSummaryResponseBody) *GetWorkTaskSummaryResponse {
	s.Body = v
	return s
}

func (s *GetWorkTaskSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
