// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlEventListResponseBody) *DescribeSdlEventListResponse
	GetBody() *DescribeSdlEventListResponseBody
}

type DescribeSdlEventListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlEventListResponse) GetBody() *DescribeSdlEventListResponseBody {
	return s.Body
}

func (s *DescribeSdlEventListResponse) SetHeaders(v map[string]*string) *DescribeSdlEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlEventListResponse) SetStatusCode(v int32) *DescribeSdlEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlEventListResponse) SetBody(v *DescribeSdlEventListResponseBody) *DescribeSdlEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
