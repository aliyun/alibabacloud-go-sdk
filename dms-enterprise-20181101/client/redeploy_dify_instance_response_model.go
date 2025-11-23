// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDifyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedeployDifyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedeployDifyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RedeployDifyInstanceResponseBody) *RedeployDifyInstanceResponse
	GetBody() *RedeployDifyInstanceResponseBody
}

type RedeployDifyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedeployDifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedeployDifyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RedeployDifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *RedeployDifyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedeployDifyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedeployDifyInstanceResponse) GetBody() *RedeployDifyInstanceResponseBody {
	return s.Body
}

func (s *RedeployDifyInstanceResponse) SetHeaders(v map[string]*string) *RedeployDifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *RedeployDifyInstanceResponse) SetStatusCode(v int32) *RedeployDifyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RedeployDifyInstanceResponse) SetBody(v *RedeployDifyInstanceResponseBody) *RedeployDifyInstanceResponse {
	s.Body = v
	return s
}

func (s *RedeployDifyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
