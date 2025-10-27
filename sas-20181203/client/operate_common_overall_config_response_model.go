// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCommonOverallConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateCommonOverallConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateCommonOverallConfigResponse
	GetStatusCode() *int32
	SetBody(v *OperateCommonOverallConfigResponseBody) *OperateCommonOverallConfigResponse
	GetBody() *OperateCommonOverallConfigResponseBody
}

type OperateCommonOverallConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateCommonOverallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateCommonOverallConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateCommonOverallConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateCommonOverallConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateCommonOverallConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateCommonOverallConfigResponse) GetBody() *OperateCommonOverallConfigResponseBody {
	return s.Body
}

func (s *OperateCommonOverallConfigResponse) SetHeaders(v map[string]*string) *OperateCommonOverallConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateCommonOverallConfigResponse) SetStatusCode(v int32) *OperateCommonOverallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateCommonOverallConfigResponse) SetBody(v *OperateCommonOverallConfigResponseBody) *OperateCommonOverallConfigResponse {
	s.Body = v
	return s
}

func (s *OperateCommonOverallConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
