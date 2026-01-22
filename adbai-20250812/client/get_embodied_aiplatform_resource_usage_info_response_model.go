// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbodiedAIPlatformResourceUsageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmbodiedAIPlatformResourceUsageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmbodiedAIPlatformResourceUsageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) *GetEmbodiedAIPlatformResourceUsageInfoResponse
	GetBody() *GetEmbodiedAIPlatformResourceUsageInfoResponseBody
}

type GetEmbodiedAIPlatformResourceUsageInfoResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmbodiedAIPlatformResourceUsageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) GetBody() *GetEmbodiedAIPlatformResourceUsageInfoResponseBody {
	return s.Body
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) SetHeaders(v map[string]*string) *GetEmbodiedAIPlatformResourceUsageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) SetStatusCode(v int32) *GetEmbodiedAIPlatformResourceUsageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) SetBody(v *GetEmbodiedAIPlatformResourceUsageInfoResponseBody) *GetEmbodiedAIPlatformResourceUsageInfoResponse {
	s.Body = v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
