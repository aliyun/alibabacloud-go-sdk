// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGranularityBySummaryIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGranularityBySummaryIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGranularityBySummaryIdResponse
	GetStatusCode() *int32
	SetBody(v *GetGranularityBySummaryIdResponseBody) *GetGranularityBySummaryIdResponse
	GetBody() *GetGranularityBySummaryIdResponseBody
}

type GetGranularityBySummaryIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGranularityBySummaryIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGranularityBySummaryIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGranularityBySummaryIdResponse) GoString() string {
	return s.String()
}

func (s *GetGranularityBySummaryIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGranularityBySummaryIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGranularityBySummaryIdResponse) GetBody() *GetGranularityBySummaryIdResponseBody {
	return s.Body
}

func (s *GetGranularityBySummaryIdResponse) SetHeaders(v map[string]*string) *GetGranularityBySummaryIdResponse {
	s.Headers = v
	return s
}

func (s *GetGranularityBySummaryIdResponse) SetStatusCode(v int32) *GetGranularityBySummaryIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGranularityBySummaryIdResponse) SetBody(v *GetGranularityBySummaryIdResponseBody) *GetGranularityBySummaryIdResponse {
	s.Body = v
	return s
}

func (s *GetGranularityBySummaryIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
