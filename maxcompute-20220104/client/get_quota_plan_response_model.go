// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaPlanResponseBody) *GetQuotaPlanResponse
	GetBody() *GetQuotaPlanResponseBody
}

type GetQuotaPlanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaPlanResponse) GetBody() *GetQuotaPlanResponseBody {
	return s.Body
}

func (s *GetQuotaPlanResponse) SetHeaders(v map[string]*string) *GetQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaPlanResponse) SetStatusCode(v int32) *GetQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaPlanResponse) SetBody(v *GetQuotaPlanResponseBody) *GetQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *GetQuotaPlanResponse) Validate() error {
	return dara.Validate(s)
}
