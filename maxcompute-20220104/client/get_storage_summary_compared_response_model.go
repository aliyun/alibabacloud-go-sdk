// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSummaryComparedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageSummaryComparedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageSummaryComparedResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageSummaryComparedResponseBody) *GetStorageSummaryComparedResponse
	GetBody() *GetStorageSummaryComparedResponseBody
}

type GetStorageSummaryComparedResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageSummaryComparedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageSummaryComparedResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSummaryComparedResponse) GoString() string {
	return s.String()
}

func (s *GetStorageSummaryComparedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageSummaryComparedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageSummaryComparedResponse) GetBody() *GetStorageSummaryComparedResponseBody {
	return s.Body
}

func (s *GetStorageSummaryComparedResponse) SetHeaders(v map[string]*string) *GetStorageSummaryComparedResponse {
	s.Headers = v
	return s
}

func (s *GetStorageSummaryComparedResponse) SetStatusCode(v int32) *GetStorageSummaryComparedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageSummaryComparedResponse) SetBody(v *GetStorageSummaryComparedResponseBody) *GetStorageSummaryComparedResponse {
	s.Body = v
	return s
}

func (s *GetStorageSummaryComparedResponse) Validate() error {
	return dara.Validate(s)
}
