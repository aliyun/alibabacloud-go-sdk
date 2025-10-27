// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageBaselineWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateImageBaselineWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateImageBaselineWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *OperateImageBaselineWhitelistResponseBody) *OperateImageBaselineWhitelistResponse
	GetBody() *OperateImageBaselineWhitelistResponseBody
}

type OperateImageBaselineWhitelistResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateImageBaselineWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateImageBaselineWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateImageBaselineWhitelistResponse) GoString() string {
	return s.String()
}

func (s *OperateImageBaselineWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateImageBaselineWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateImageBaselineWhitelistResponse) GetBody() *OperateImageBaselineWhitelistResponseBody {
	return s.Body
}

func (s *OperateImageBaselineWhitelistResponse) SetHeaders(v map[string]*string) *OperateImageBaselineWhitelistResponse {
	s.Headers = v
	return s
}

func (s *OperateImageBaselineWhitelistResponse) SetStatusCode(v int32) *OperateImageBaselineWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateImageBaselineWhitelistResponse) SetBody(v *OperateImageBaselineWhitelistResponseBody) *OperateImageBaselineWhitelistResponse {
	s.Body = v
	return s
}

func (s *OperateImageBaselineWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
