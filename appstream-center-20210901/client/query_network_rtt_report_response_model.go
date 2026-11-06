// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetworkRttReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNetworkRttReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNetworkRttReportResponse
	GetStatusCode() *int32
	SetBody(v *QueryNetworkRttReportResponseBody) *QueryNetworkRttReportResponse
	GetBody() *QueryNetworkRttReportResponseBody
}

type QueryNetworkRttReportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNetworkRttReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNetworkRttReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNetworkRttReportResponse) GoString() string {
	return s.String()
}

func (s *QueryNetworkRttReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNetworkRttReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNetworkRttReportResponse) GetBody() *QueryNetworkRttReportResponseBody {
	return s.Body
}

func (s *QueryNetworkRttReportResponse) SetHeaders(v map[string]*string) *QueryNetworkRttReportResponse {
	s.Headers = v
	return s
}

func (s *QueryNetworkRttReportResponse) SetStatusCode(v int32) *QueryNetworkRttReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNetworkRttReportResponse) SetBody(v *QueryNetworkRttReportResponseBody) *QueryNetworkRttReportResponse {
	s.Body = v
	return s
}

func (s *QueryNetworkRttReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
