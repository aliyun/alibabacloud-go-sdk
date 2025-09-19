// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspEventExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspEventExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspEventExportInfoResponseBody) *DescribeSuspEventExportInfoResponse
	GetBody() *DescribeSuspEventExportInfoResponseBody
}

type DescribeSuspEventExportInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspEventExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspEventExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspEventExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspEventExportInfoResponse) GetBody() *DescribeSuspEventExportInfoResponseBody {
	return s.Body
}

func (s *DescribeSuspEventExportInfoResponse) SetHeaders(v map[string]*string) *DescribeSuspEventExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventExportInfoResponse) SetStatusCode(v int32) *DescribeSuspEventExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponse) SetBody(v *DescribeSuspEventExportInfoResponseBody) *DescribeSuspEventExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspEventExportInfoResponse) Validate() error {
	return dara.Validate(s)
}
