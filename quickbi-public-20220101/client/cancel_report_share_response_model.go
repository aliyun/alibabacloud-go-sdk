// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelReportShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelReportShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelReportShareResponse
	GetStatusCode() *int32
	SetBody(v *CancelReportShareResponseBody) *CancelReportShareResponse
	GetBody() *CancelReportShareResponseBody
}

type CancelReportShareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelReportShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelReportShareResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelReportShareResponse) GoString() string {
	return s.String()
}

func (s *CancelReportShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelReportShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelReportShareResponse) GetBody() *CancelReportShareResponseBody {
	return s.Body
}

func (s *CancelReportShareResponse) SetHeaders(v map[string]*string) *CancelReportShareResponse {
	s.Headers = v
	return s
}

func (s *CancelReportShareResponse) SetStatusCode(v int32) *CancelReportShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelReportShareResponse) SetBody(v *CancelReportShareResponseBody) *CancelReportShareResponse {
	s.Body = v
	return s
}

func (s *CancelReportShareResponse) Validate() error {
	return dara.Validate(s)
}
