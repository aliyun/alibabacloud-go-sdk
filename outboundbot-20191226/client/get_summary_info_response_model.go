// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSummaryInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSummaryInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSummaryInfoResponseBody) *GetSummaryInfoResponse
	GetBody() *GetSummaryInfoResponseBody
}

type GetSummaryInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSummaryInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSummaryInfoResponse) GetBody() *GetSummaryInfoResponseBody {
	return s.Body
}

func (s *GetSummaryInfoResponse) SetHeaders(v map[string]*string) *GetSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryInfoResponse) SetStatusCode(v int32) *GetSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryInfoResponse) SetBody(v *GetSummaryInfoResponseBody) *GetSummaryInfoResponse {
	s.Body = v
	return s
}

func (s *GetSummaryInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
