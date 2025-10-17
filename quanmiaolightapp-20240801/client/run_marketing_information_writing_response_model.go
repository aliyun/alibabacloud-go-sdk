// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunMarketingInformationWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunMarketingInformationWritingResponse
	GetStatusCode() *int32
	SetBody(v *RunMarketingInformationWritingResponseBody) *RunMarketingInformationWritingResponse
	GetBody() *RunMarketingInformationWritingResponseBody
}

type RunMarketingInformationWritingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunMarketingInformationWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMarketingInformationWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponse) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunMarketingInformationWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunMarketingInformationWritingResponse) GetBody() *RunMarketingInformationWritingResponseBody {
	return s.Body
}

func (s *RunMarketingInformationWritingResponse) SetHeaders(v map[string]*string) *RunMarketingInformationWritingResponse {
	s.Headers = v
	return s
}

func (s *RunMarketingInformationWritingResponse) SetStatusCode(v int32) *RunMarketingInformationWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMarketingInformationWritingResponse) SetBody(v *RunMarketingInformationWritingResponseBody) *RunMarketingInformationWritingResponse {
	s.Body = v
	return s
}

func (s *RunMarketingInformationWritingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
