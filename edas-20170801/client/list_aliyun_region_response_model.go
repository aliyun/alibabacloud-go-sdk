// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliyunRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliyunRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliyunRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListAliyunRegionResponseBody) *ListAliyunRegionResponse
	GetBody() *ListAliyunRegionResponseBody
}

type ListAliyunRegionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliyunRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliyunRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunRegionResponse) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliyunRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliyunRegionResponse) GetBody() *ListAliyunRegionResponseBody {
	return s.Body
}

func (s *ListAliyunRegionResponse) SetHeaders(v map[string]*string) *ListAliyunRegionResponse {
	s.Headers = v
	return s
}

func (s *ListAliyunRegionResponse) SetStatusCode(v int32) *ListAliyunRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliyunRegionResponse) SetBody(v *ListAliyunRegionResponseBody) *ListAliyunRegionResponse {
	s.Body = v
	return s
}

func (s *ListAliyunRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
