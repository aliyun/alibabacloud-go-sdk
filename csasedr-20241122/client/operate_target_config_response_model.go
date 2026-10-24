// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *OperateTargetConfigResponseBody) *OperateTargetConfigResponse
	GetBody() *OperateTargetConfigResponseBody
}

type OperateTargetConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateTargetConfigResponse) GetBody() *OperateTargetConfigResponseBody {
	return s.Body
}

func (s *OperateTargetConfigResponse) SetHeaders(v map[string]*string) *OperateTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateTargetConfigResponse) SetStatusCode(v int32) *OperateTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateTargetConfigResponse) SetBody(v *OperateTargetConfigResponseBody) *OperateTargetConfigResponse {
	s.Body = v
	return s
}

func (s *OperateTargetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
