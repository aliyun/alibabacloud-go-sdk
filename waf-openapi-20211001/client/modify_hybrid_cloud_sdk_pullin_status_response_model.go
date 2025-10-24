// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudSdkPullinStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudSdkPullinStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudSdkPullinStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudSdkPullinStatusResponseBody) *ModifyHybridCloudSdkPullinStatusResponse
	GetBody() *ModifyHybridCloudSdkPullinStatusResponseBody
}

type ModifyHybridCloudSdkPullinStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudSdkPullinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudSdkPullinStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudSdkPullinStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) GetBody() *ModifyHybridCloudSdkPullinStatusResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudSdkPullinStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) SetStatusCode(v int32) *ModifyHybridCloudSdkPullinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) SetBody(v *ModifyHybridCloudSdkPullinStatusResponseBody) *ModifyHybridCloudSdkPullinStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
