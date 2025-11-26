// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAutoRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterAutoRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterAutoRenewResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterAutoRenewResponseBody) *UpdateClusterAutoRenewResponse
	GetBody() *UpdateClusterAutoRenewResponseBody
}

type UpdateClusterAutoRenewResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterAutoRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterAutoRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAutoRenewResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterAutoRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterAutoRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterAutoRenewResponse) GetBody() *UpdateClusterAutoRenewResponseBody {
	return s.Body
}

func (s *UpdateClusterAutoRenewResponse) SetHeaders(v map[string]*string) *UpdateClusterAutoRenewResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterAutoRenewResponse) SetStatusCode(v int32) *UpdateClusterAutoRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterAutoRenewResponse) SetBody(v *UpdateClusterAutoRenewResponseBody) *UpdateClusterAutoRenewResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterAutoRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
