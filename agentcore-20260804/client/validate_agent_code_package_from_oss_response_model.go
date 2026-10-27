// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAgentCodePackageFromOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateAgentCodePackageFromOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateAgentCodePackageFromOssResponse
	GetStatusCode() *int32
	SetBody(v *ValidateAgentCodePackageFromOssResponseBody) *ValidateAgentCodePackageFromOssResponse
	GetBody() *ValidateAgentCodePackageFromOssResponseBody
}

type ValidateAgentCodePackageFromOssResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateAgentCodePackageFromOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateAgentCodePackageFromOssResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssResponse) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateAgentCodePackageFromOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateAgentCodePackageFromOssResponse) GetBody() *ValidateAgentCodePackageFromOssResponseBody {
	return s.Body
}

func (s *ValidateAgentCodePackageFromOssResponse) SetHeaders(v map[string]*string) *ValidateAgentCodePackageFromOssResponse {
	s.Headers = v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponse) SetStatusCode(v int32) *ValidateAgentCodePackageFromOssResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponse) SetBody(v *ValidateAgentCodePackageFromOssResponseBody) *ValidateAgentCodePackageFromOssResponse {
	s.Body = v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
