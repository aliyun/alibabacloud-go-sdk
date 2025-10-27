// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateCommonOverallConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchOperateCommonOverallConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchOperateCommonOverallConfigResponse
	GetStatusCode() *int32
	SetBody(v *BatchOperateCommonOverallConfigResponseBody) *BatchOperateCommonOverallConfigResponse
	GetBody() *BatchOperateCommonOverallConfigResponseBody
}

type BatchOperateCommonOverallConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchOperateCommonOverallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchOperateCommonOverallConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateCommonOverallConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchOperateCommonOverallConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchOperateCommonOverallConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchOperateCommonOverallConfigResponse) GetBody() *BatchOperateCommonOverallConfigResponseBody {
	return s.Body
}

func (s *BatchOperateCommonOverallConfigResponse) SetHeaders(v map[string]*string) *BatchOperateCommonOverallConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchOperateCommonOverallConfigResponse) SetStatusCode(v int32) *BatchOperateCommonOverallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchOperateCommonOverallConfigResponse) SetBody(v *BatchOperateCommonOverallConfigResponseBody) *BatchOperateCommonOverallConfigResponse {
	s.Body = v
	return s
}

func (s *BatchOperateCommonOverallConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
