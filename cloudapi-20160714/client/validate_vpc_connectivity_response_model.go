// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateVpcConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateVpcConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateVpcConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *ValidateVpcConnectivityResponseBody) *ValidateVpcConnectivityResponse
	GetBody() *ValidateVpcConnectivityResponseBody
}

type ValidateVpcConnectivityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateVpcConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateVpcConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateVpcConnectivityResponse) GoString() string {
	return s.String()
}

func (s *ValidateVpcConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateVpcConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateVpcConnectivityResponse) GetBody() *ValidateVpcConnectivityResponseBody {
	return s.Body
}

func (s *ValidateVpcConnectivityResponse) SetHeaders(v map[string]*string) *ValidateVpcConnectivityResponse {
	s.Headers = v
	return s
}

func (s *ValidateVpcConnectivityResponse) SetStatusCode(v int32) *ValidateVpcConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateVpcConnectivityResponse) SetBody(v *ValidateVpcConnectivityResponseBody) *ValidateVpcConnectivityResponse {
	s.Body = v
	return s
}

func (s *ValidateVpcConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
