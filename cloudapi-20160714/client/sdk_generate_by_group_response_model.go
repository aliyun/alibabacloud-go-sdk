// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdkGenerateByGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdkGenerateByGroupResponse
	GetStatusCode() *int32
	SetBody(v *SdkGenerateByGroupResponseBody) *SdkGenerateByGroupResponse
	GetBody() *SdkGenerateByGroupResponseBody
}

type SdkGenerateByGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdkGenerateByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdkGenerateByGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByGroupResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdkGenerateByGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdkGenerateByGroupResponse) GetBody() *SdkGenerateByGroupResponseBody {
	return s.Body
}

func (s *SdkGenerateByGroupResponse) SetHeaders(v map[string]*string) *SdkGenerateByGroupResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateByGroupResponse) SetStatusCode(v int32) *SdkGenerateByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateByGroupResponse) SetBody(v *SdkGenerateByGroupResponseBody) *SdkGenerateByGroupResponse {
	s.Body = v
	return s
}

func (s *SdkGenerateByGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
