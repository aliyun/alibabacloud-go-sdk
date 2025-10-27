// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthSummaryResponseBody) *GetAuthSummaryResponse
	GetBody() *GetAuthSummaryResponseBody
}

type GetAuthSummaryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAuthSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthSummaryResponse) GetBody() *GetAuthSummaryResponseBody {
	return s.Body
}

func (s *GetAuthSummaryResponse) SetHeaders(v map[string]*string) *GetAuthSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAuthSummaryResponse) SetStatusCode(v int32) *GetAuthSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthSummaryResponse) SetBody(v *GetAuthSummaryResponseBody) *GetAuthSummaryResponse {
	s.Body = v
	return s
}

func (s *GetAuthSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
