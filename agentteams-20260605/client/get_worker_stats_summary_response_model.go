// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerStatsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkerStatsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkerStatsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkerStatsSummaryResponseBody) *GetWorkerStatsSummaryResponse
	GetBody() *GetWorkerStatsSummaryResponseBody
}

type GetWorkerStatsSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerStatsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkerStatsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerStatsSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerStatsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkerStatsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkerStatsSummaryResponse) GetBody() *GetWorkerStatsSummaryResponseBody {
	return s.Body
}

func (s *GetWorkerStatsSummaryResponse) SetHeaders(v map[string]*string) *GetWorkerStatsSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerStatsSummaryResponse) SetStatusCode(v int32) *GetWorkerStatsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerStatsSummaryResponse) SetBody(v *GetWorkerStatsSummaryResponseBody) *GetWorkerStatsSummaryResponse {
	s.Body = v
	return s
}

func (s *GetWorkerStatsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
