// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIvrTrackingSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIvrTrackingSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIvrTrackingSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetIvrTrackingSummaryResponseBody) *GetIvrTrackingSummaryResponse
	GetBody() *GetIvrTrackingSummaryResponseBody
}

type GetIvrTrackingSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIvrTrackingSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIvrTrackingSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIvrTrackingSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetIvrTrackingSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIvrTrackingSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIvrTrackingSummaryResponse) GetBody() *GetIvrTrackingSummaryResponseBody {
	return s.Body
}

func (s *GetIvrTrackingSummaryResponse) SetHeaders(v map[string]*string) *GetIvrTrackingSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetIvrTrackingSummaryResponse) SetStatusCode(v int32) *GetIvrTrackingSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIvrTrackingSummaryResponse) SetBody(v *GetIvrTrackingSummaryResponseBody) *GetIvrTrackingSummaryResponse {
	s.Body = v
	return s
}

func (s *GetIvrTrackingSummaryResponse) Validate() error {
	return dara.Validate(s)
}
