// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseUserResponseBody) *DescribeLangfuseUserResponse
	GetBody() *DescribeLangfuseUserResponseBody
}

type DescribeLangfuseUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseUserResponse) GetBody() *DescribeLangfuseUserResponseBody {
	return s.Body
}

func (s *DescribeLangfuseUserResponse) SetHeaders(v map[string]*string) *DescribeLangfuseUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseUserResponse) SetStatusCode(v int32) *DescribeLangfuseUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseUserResponse) SetBody(v *DescribeLangfuseUserResponseBody) *DescribeLangfuseUserResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
