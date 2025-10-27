// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCommonTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateCommonTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateCommonTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *OperateCommonTargetConfigResponseBody) *OperateCommonTargetConfigResponse
	GetBody() *OperateCommonTargetConfigResponseBody
}

type OperateCommonTargetConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateCommonTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateCommonTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateCommonTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateCommonTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateCommonTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateCommonTargetConfigResponse) GetBody() *OperateCommonTargetConfigResponseBody {
	return s.Body
}

func (s *OperateCommonTargetConfigResponse) SetHeaders(v map[string]*string) *OperateCommonTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateCommonTargetConfigResponse) SetStatusCode(v int32) *OperateCommonTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateCommonTargetConfigResponse) SetBody(v *OperateCommonTargetConfigResponseBody) *OperateCommonTargetConfigResponse {
	s.Body = v
	return s
}

func (s *OperateCommonTargetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
