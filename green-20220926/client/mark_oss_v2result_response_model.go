// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkOssV2ResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MarkOssV2ResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MarkOssV2ResultResponse
	GetStatusCode() *int32
	SetBody(v *MarkOssV2ResultResponseBody) *MarkOssV2ResultResponse
	GetBody() *MarkOssV2ResultResponseBody
}

type MarkOssV2ResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkOssV2ResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkOssV2ResultResponse) String() string {
	return dara.Prettify(s)
}

func (s MarkOssV2ResultResponse) GoString() string {
	return s.String()
}

func (s *MarkOssV2ResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MarkOssV2ResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MarkOssV2ResultResponse) GetBody() *MarkOssV2ResultResponseBody {
	return s.Body
}

func (s *MarkOssV2ResultResponse) SetHeaders(v map[string]*string) *MarkOssV2ResultResponse {
	s.Headers = v
	return s
}

func (s *MarkOssV2ResultResponse) SetStatusCode(v int32) *MarkOssV2ResultResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkOssV2ResultResponse) SetBody(v *MarkOssV2ResultResponseBody) *MarkOssV2ResultResponse {
	s.Body = v
	return s
}

func (s *MarkOssV2ResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
