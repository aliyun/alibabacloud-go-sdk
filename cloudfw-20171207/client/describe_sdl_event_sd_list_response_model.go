// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventSdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlEventSdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlEventSdListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlEventSdListResponseBody) *DescribeSdlEventSdListResponse
	GetBody() *DescribeSdlEventSdListResponseBody
}

type DescribeSdlEventSdListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlEventSdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlEventSdListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventSdListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventSdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlEventSdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlEventSdListResponse) GetBody() *DescribeSdlEventSdListResponseBody {
	return s.Body
}

func (s *DescribeSdlEventSdListResponse) SetHeaders(v map[string]*string) *DescribeSdlEventSdListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlEventSdListResponse) SetStatusCode(v int32) *DescribeSdlEventSdListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlEventSdListResponse) SetBody(v *DescribeSdlEventSdListResponseBody) *DescribeSdlEventSdListResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlEventSdListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
