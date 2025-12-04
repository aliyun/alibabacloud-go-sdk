// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestartInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRestartInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRestartInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CancelRestartInstanceResponseBody) *CancelRestartInstanceResponse
	GetBody() *CancelRestartInstanceResponseBody
}

type CancelRestartInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRestartInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *CancelRestartInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRestartInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRestartInstanceResponse) GetBody() *CancelRestartInstanceResponseBody {
	return s.Body
}

func (s *CancelRestartInstanceResponse) SetHeaders(v map[string]*string) *CancelRestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *CancelRestartInstanceResponse) SetStatusCode(v int32) *CancelRestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRestartInstanceResponse) SetBody(v *CancelRestartInstanceResponseBody) *CancelRestartInstanceResponse {
	s.Body = v
	return s
}

func (s *CancelRestartInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
