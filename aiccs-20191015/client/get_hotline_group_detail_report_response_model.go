// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineGroupDetailReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineGroupDetailReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineGroupDetailReportResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineGroupDetailReportResponseBody) *GetHotlineGroupDetailReportResponse
	GetBody() *GetHotlineGroupDetailReportResponseBody
}

type GetHotlineGroupDetailReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineGroupDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineGroupDetailReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineGroupDetailReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineGroupDetailReportResponse) GetBody() *GetHotlineGroupDetailReportResponseBody {
	return s.Body
}

func (s *GetHotlineGroupDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineGroupDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetStatusCode(v int32) *GetHotlineGroupDetailReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetBody(v *GetHotlineGroupDetailReportResponseBody) *GetHotlineGroupDetailReportResponse {
	s.Body = v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
