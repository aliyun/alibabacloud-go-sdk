// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuspEventSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuspEventSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetSuspEventSummaryResponseBody) *GetSuspEventSummaryResponse
	GetBody() *GetSuspEventSummaryResponseBody
}

type GetSuspEventSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuspEventSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuspEventSummaryResponse) GetBody() *GetSuspEventSummaryResponseBody {
	return s.Body
}

func (s *GetSuspEventSummaryResponse) SetHeaders(v map[string]*string) *GetSuspEventSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventSummaryResponse) SetStatusCode(v int32) *GetSuspEventSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponse) SetBody(v *GetSuspEventSummaryResponseBody) *GetSuspEventSummaryResponse {
	s.Body = v
	return s
}

func (s *GetSuspEventSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
