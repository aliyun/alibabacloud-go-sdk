// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveCountSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResolveCountSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResolveCountSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetResolveCountSummaryResponseBody) *GetResolveCountSummaryResponse
	GetBody() *GetResolveCountSummaryResponseBody
}

type GetResolveCountSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResolveCountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResolveCountSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResolveCountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResolveCountSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResolveCountSummaryResponse) GetBody() *GetResolveCountSummaryResponseBody {
	return s.Body
}

func (s *GetResolveCountSummaryResponse) SetHeaders(v map[string]*string) *GetResolveCountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetResolveCountSummaryResponse) SetStatusCode(v int32) *GetResolveCountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResolveCountSummaryResponse) SetBody(v *GetResolveCountSummaryResponseBody) *GetResolveCountSummaryResponse {
	s.Body = v
	return s
}

func (s *GetResolveCountSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
