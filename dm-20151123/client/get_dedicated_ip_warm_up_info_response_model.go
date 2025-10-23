// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDedicatedIpWarmUpInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDedicatedIpWarmUpInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDedicatedIpWarmUpInfoResponseBody) *GetDedicatedIpWarmUpInfoResponse
	GetBody() *GetDedicatedIpWarmUpInfoResponseBody
}

type GetDedicatedIpWarmUpInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDedicatedIpWarmUpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDedicatedIpWarmUpInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDedicatedIpWarmUpInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDedicatedIpWarmUpInfoResponse) GetBody() *GetDedicatedIpWarmUpInfoResponseBody {
	return s.Body
}

func (s *GetDedicatedIpWarmUpInfoResponse) SetHeaders(v map[string]*string) *GetDedicatedIpWarmUpInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponse) SetStatusCode(v int32) *GetDedicatedIpWarmUpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponse) SetBody(v *GetDedicatedIpWarmUpInfoResponseBody) *GetDedicatedIpWarmUpInfoResponse {
	s.Body = v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
