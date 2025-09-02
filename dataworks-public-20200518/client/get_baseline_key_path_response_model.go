// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineKeyPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaselineKeyPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaselineKeyPathResponse
	GetStatusCode() *int32
	SetBody(v *GetBaselineKeyPathResponseBody) *GetBaselineKeyPathResponse
	GetBody() *GetBaselineKeyPathResponseBody
}

type GetBaselineKeyPathResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineKeyPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineKeyPathResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaselineKeyPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaselineKeyPathResponse) GetBody() *GetBaselineKeyPathResponseBody {
	return s.Body
}

func (s *GetBaselineKeyPathResponse) SetHeaders(v map[string]*string) *GetBaselineKeyPathResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineKeyPathResponse) SetStatusCode(v int32) *GetBaselineKeyPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineKeyPathResponse) SetBody(v *GetBaselineKeyPathResponseBody) *GetBaselineKeyPathResponse {
	s.Body = v
	return s
}

func (s *GetBaselineKeyPathResponse) Validate() error {
	return dara.Validate(s)
}
