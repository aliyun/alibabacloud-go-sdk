// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionMemberAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateYikeProductionMemberAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateYikeProductionMemberAuthResponse
	GetStatusCode() *int32
	SetBody(v *UpdateYikeProductionMemberAuthResponseBody) *UpdateYikeProductionMemberAuthResponse
	GetBody() *UpdateYikeProductionMemberAuthResponseBody
}

type UpdateYikeProductionMemberAuthResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateYikeProductionMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateYikeProductionMemberAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionMemberAuthResponse) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionMemberAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateYikeProductionMemberAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateYikeProductionMemberAuthResponse) GetBody() *UpdateYikeProductionMemberAuthResponseBody {
	return s.Body
}

func (s *UpdateYikeProductionMemberAuthResponse) SetHeaders(v map[string]*string) *UpdateYikeProductionMemberAuthResponse {
	s.Headers = v
	return s
}

func (s *UpdateYikeProductionMemberAuthResponse) SetStatusCode(v int32) *UpdateYikeProductionMemberAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthResponse) SetBody(v *UpdateYikeProductionMemberAuthResponseBody) *UpdateYikeProductionMemberAuthResponse {
	s.Body = v
	return s
}

func (s *UpdateYikeProductionMemberAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
