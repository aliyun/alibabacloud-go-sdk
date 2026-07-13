// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskStatsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskStatsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskStatsSummaryResponseBody) *GetTaskStatsSummaryResponse
	GetBody() *GetTaskStatsSummaryResponseBody
}

type GetTaskStatsSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatsSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskStatsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskStatsSummaryResponse) GetBody() *GetTaskStatsSummaryResponseBody {
	return s.Body
}

func (s *GetTaskStatsSummaryResponse) SetHeaders(v map[string]*string) *GetTaskStatsSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatsSummaryResponse) SetStatusCode(v int32) *GetTaskStatsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatsSummaryResponse) SetBody(v *GetTaskStatsSummaryResponseBody) *GetTaskStatsSummaryResponse {
	s.Body = v
	return s
}

func (s *GetTaskStatsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
