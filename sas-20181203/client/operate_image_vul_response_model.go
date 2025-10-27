// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageVulResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateImageVulResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateImageVulResponse
	GetStatusCode() *int32
	SetBody(v *OperateImageVulResponseBody) *OperateImageVulResponse
	GetBody() *OperateImageVulResponseBody
}

type OperateImageVulResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateImageVulResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateImageVulResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateImageVulResponse) GoString() string {
	return s.String()
}

func (s *OperateImageVulResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateImageVulResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateImageVulResponse) GetBody() *OperateImageVulResponseBody {
	return s.Body
}

func (s *OperateImageVulResponse) SetHeaders(v map[string]*string) *OperateImageVulResponse {
	s.Headers = v
	return s
}

func (s *OperateImageVulResponse) SetStatusCode(v int32) *OperateImageVulResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateImageVulResponse) SetBody(v *OperateImageVulResponseBody) *OperateImageVulResponse {
	s.Body = v
	return s
}

func (s *OperateImageVulResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
