// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspEventUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspEventUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspEventUserSettingResponseBody) *DescribeSuspEventUserSettingResponse
	GetBody() *DescribeSuspEventUserSettingResponseBody
}

type DescribeSuspEventUserSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspEventUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspEventUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventUserSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspEventUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspEventUserSettingResponse) GetBody() *DescribeSuspEventUserSettingResponseBody {
	return s.Body
}

func (s *DescribeSuspEventUserSettingResponse) SetHeaders(v map[string]*string) *DescribeSuspEventUserSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventUserSettingResponse) SetStatusCode(v int32) *DescribeSuspEventUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspEventUserSettingResponse) SetBody(v *DescribeSuspEventUserSettingResponseBody) *DescribeSuspEventUserSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspEventUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
