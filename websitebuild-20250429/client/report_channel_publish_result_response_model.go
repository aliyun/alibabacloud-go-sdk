// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportChannelPublishResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportChannelPublishResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportChannelPublishResultResponse
	GetStatusCode() *int32
	SetBody(v *ReportChannelPublishResultResponseBody) *ReportChannelPublishResultResponse
	GetBody() *ReportChannelPublishResultResponseBody
}

type ReportChannelPublishResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportChannelPublishResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportChannelPublishResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportChannelPublishResultResponse) GoString() string {
	return s.String()
}

func (s *ReportChannelPublishResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportChannelPublishResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportChannelPublishResultResponse) GetBody() *ReportChannelPublishResultResponseBody {
	return s.Body
}

func (s *ReportChannelPublishResultResponse) SetHeaders(v map[string]*string) *ReportChannelPublishResultResponse {
	s.Headers = v
	return s
}

func (s *ReportChannelPublishResultResponse) SetStatusCode(v int32) *ReportChannelPublishResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportChannelPublishResultResponse) SetBody(v *ReportChannelPublishResultResponseBody) *ReportChannelPublishResultResponse {
	s.Body = v
	return s
}

func (s *ReportChannelPublishResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
