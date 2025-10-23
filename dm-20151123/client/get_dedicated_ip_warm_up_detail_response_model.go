// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDedicatedIpWarmUpDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDedicatedIpWarmUpDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDedicatedIpWarmUpDetailResponseBody) *GetDedicatedIpWarmUpDetailResponse
	GetBody() *GetDedicatedIpWarmUpDetailResponseBody
}

type GetDedicatedIpWarmUpDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDedicatedIpWarmUpDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDedicatedIpWarmUpDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDedicatedIpWarmUpDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDedicatedIpWarmUpDetailResponse) GetBody() *GetDedicatedIpWarmUpDetailResponseBody {
	return s.Body
}

func (s *GetDedicatedIpWarmUpDetailResponse) SetHeaders(v map[string]*string) *GetDedicatedIpWarmUpDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponse) SetStatusCode(v int32) *GetDedicatedIpWarmUpDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponse) SetBody(v *GetDedicatedIpWarmUpDetailResponseBody) *GetDedicatedIpWarmUpDetailResponse {
	s.Body = v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
