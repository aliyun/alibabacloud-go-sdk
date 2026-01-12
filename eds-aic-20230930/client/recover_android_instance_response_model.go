// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RecoverAndroidInstanceResponseBody) *RecoverAndroidInstanceResponse
	GetBody() *RecoverAndroidInstanceResponseBody
}

type RecoverAndroidInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *RecoverAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverAndroidInstanceResponse) GetBody() *RecoverAndroidInstanceResponseBody {
	return s.Body
}

func (s *RecoverAndroidInstanceResponse) SetHeaders(v map[string]*string) *RecoverAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *RecoverAndroidInstanceResponse) SetStatusCode(v int32) *RecoverAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverAndroidInstanceResponse) SetBody(v *RecoverAndroidInstanceResponseBody) *RecoverAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *RecoverAndroidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
