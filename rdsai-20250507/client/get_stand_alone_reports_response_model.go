// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandAloneReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandAloneReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandAloneReportsResponse
	GetStatusCode() *int32
	SetBody(v *GetStandAloneReportsResponseBody) *GetStandAloneReportsResponse
	GetBody() *GetStandAloneReportsResponseBody
}

type GetStandAloneReportsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandAloneReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandAloneReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandAloneReportsResponse) GoString() string {
	return s.String()
}

func (s *GetStandAloneReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandAloneReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandAloneReportsResponse) GetBody() *GetStandAloneReportsResponseBody {
	return s.Body
}

func (s *GetStandAloneReportsResponse) SetHeaders(v map[string]*string) *GetStandAloneReportsResponse {
	s.Headers = v
	return s
}

func (s *GetStandAloneReportsResponse) SetStatusCode(v int32) *GetStandAloneReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandAloneReportsResponse) SetBody(v *GetStandAloneReportsResponseBody) *GetStandAloneReportsResponse {
	s.Body = v
	return s
}

func (s *GetStandAloneReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
