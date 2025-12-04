// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckModifyConfigNeedRestartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckModifyConfigNeedRestartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckModifyConfigNeedRestartResponse
	GetStatusCode() *int32
	SetBody(v *CheckModifyConfigNeedRestartResponseBody) *CheckModifyConfigNeedRestartResponse
	GetBody() *CheckModifyConfigNeedRestartResponseBody
}

type CheckModifyConfigNeedRestartResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckModifyConfigNeedRestartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckModifyConfigNeedRestartResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckModifyConfigNeedRestartResponse) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckModifyConfigNeedRestartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckModifyConfigNeedRestartResponse) GetBody() *CheckModifyConfigNeedRestartResponseBody {
	return s.Body
}

func (s *CheckModifyConfigNeedRestartResponse) SetHeaders(v map[string]*string) *CheckModifyConfigNeedRestartResponse {
	s.Headers = v
	return s
}

func (s *CheckModifyConfigNeedRestartResponse) SetStatusCode(v int32) *CheckModifyConfigNeedRestartResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckModifyConfigNeedRestartResponse) SetBody(v *CheckModifyConfigNeedRestartResponseBody) *CheckModifyConfigNeedRestartResponse {
	s.Body = v
	return s
}

func (s *CheckModifyConfigNeedRestartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
