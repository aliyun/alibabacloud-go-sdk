// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckScopeConfigInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCheckScopeConfigInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCheckScopeConfigInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCheckScopeConfigInstanceResponseBody) *ChangeCheckScopeConfigInstanceResponse
	GetBody() *ChangeCheckScopeConfigInstanceResponseBody
}

type ChangeCheckScopeConfigInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCheckScopeConfigInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCheckScopeConfigInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckScopeConfigInstanceResponse) GoString() string {
	return s.String()
}

func (s *ChangeCheckScopeConfigInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCheckScopeConfigInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCheckScopeConfigInstanceResponse) GetBody() *ChangeCheckScopeConfigInstanceResponseBody {
	return s.Body
}

func (s *ChangeCheckScopeConfigInstanceResponse) SetHeaders(v map[string]*string) *ChangeCheckScopeConfigInstanceResponse {
	s.Headers = v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponse) SetStatusCode(v int32) *ChangeCheckScopeConfigInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponse) SetBody(v *ChangeCheckScopeConfigInstanceResponseBody) *ChangeCheckScopeConfigInstanceResponse {
	s.Body = v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
