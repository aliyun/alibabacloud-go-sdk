// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationExtractResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunMarketingInformationExtractResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunMarketingInformationExtractResponse
	GetStatusCode() *int32
	SetId(v string) *RunMarketingInformationExtractResponse
	GetId() *string
	SetEvent(v string) *RunMarketingInformationExtractResponse
	GetEvent() *string
	SetBody(v *RunMarketingInformationExtractResponseBody) *RunMarketingInformationExtractResponse
	GetBody() *RunMarketingInformationExtractResponseBody
}

type RunMarketingInformationExtractResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunMarketingInformationExtractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMarketingInformationExtractResponse) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponse) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunMarketingInformationExtractResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunMarketingInformationExtractResponse) GetId() *string {
	return s.Id
}

func (s *RunMarketingInformationExtractResponse) GetEvent() *string {
	return s.Event
}

func (s *RunMarketingInformationExtractResponse) GetBody() *RunMarketingInformationExtractResponseBody {
	return s.Body
}

func (s *RunMarketingInformationExtractResponse) SetHeaders(v map[string]*string) *RunMarketingInformationExtractResponse {
	s.Headers = v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetStatusCode(v int32) *RunMarketingInformationExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetId(v string) *RunMarketingInformationExtractResponse {
	s.Id = &v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetEvent(v string) *RunMarketingInformationExtractResponse {
	s.Event = &v
	return s
}

func (s *RunMarketingInformationExtractResponse) SetBody(v *RunMarketingInformationExtractResponseBody) *RunMarketingInformationExtractResponse {
	s.Body = v
	return s
}

func (s *RunMarketingInformationExtractResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
