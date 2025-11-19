// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscodeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscodeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscodeSummaryResponseBody) *GetTranscodeSummaryResponse
	GetBody() *GetTranscodeSummaryResponseBody
}

type GetTranscodeSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscodeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscodeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscodeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscodeSummaryResponse) GetBody() *GetTranscodeSummaryResponseBody {
	return s.Body
}

func (s *GetTranscodeSummaryResponse) SetHeaders(v map[string]*string) *GetTranscodeSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeSummaryResponse) SetStatusCode(v int32) *GetTranscodeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscodeSummaryResponse) SetBody(v *GetTranscodeSummaryResponseBody) *GetTranscodeSummaryResponse {
	s.Body = v
	return s
}

func (s *GetTranscodeSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
