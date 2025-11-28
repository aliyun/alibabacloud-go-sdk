// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBizChainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBizChainResponseBody) *UpdateBizChainResponse
	GetBody() *UpdateBizChainResponseBody
}

type UpdateBizChainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizChainResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBizChainResponse) GetBody() *UpdateBizChainResponseBody {
	return s.Body
}

func (s *UpdateBizChainResponse) SetHeaders(v map[string]*string) *UpdateBizChainResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizChainResponse) SetStatusCode(v int32) *UpdateBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizChainResponse) SetBody(v *UpdateBizChainResponseBody) *UpdateBizChainResponse {
	s.Body = v
	return s
}

func (s *UpdateBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
