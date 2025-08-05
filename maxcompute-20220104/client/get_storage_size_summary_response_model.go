// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSizeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageSizeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageSizeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageSizeSummaryResponseBody) *GetStorageSizeSummaryResponse
	GetBody() *GetStorageSizeSummaryResponseBody
}

type GetStorageSizeSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageSizeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageSizeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSizeSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetStorageSizeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageSizeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageSizeSummaryResponse) GetBody() *GetStorageSizeSummaryResponseBody {
	return s.Body
}

func (s *GetStorageSizeSummaryResponse) SetHeaders(v map[string]*string) *GetStorageSizeSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetStorageSizeSummaryResponse) SetStatusCode(v int32) *GetStorageSizeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageSizeSummaryResponse) SetBody(v *GetStorageSizeSummaryResponseBody) *GetStorageSizeSummaryResponse {
	s.Body = v
	return s
}

func (s *GetStorageSizeSummaryResponse) Validate() error {
	return dara.Validate(s)
}
