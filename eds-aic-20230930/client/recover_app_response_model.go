// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverAppResponse
	GetStatusCode() *int32
	SetBody(v *RecoverAppResponseBody) *RecoverAppResponse
	GetBody() *RecoverAppResponseBody
}

type RecoverAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverAppResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppResponse) GoString() string {
	return s.String()
}

func (s *RecoverAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverAppResponse) GetBody() *RecoverAppResponseBody {
	return s.Body
}

func (s *RecoverAppResponse) SetHeaders(v map[string]*string) *RecoverAppResponse {
	s.Headers = v
	return s
}

func (s *RecoverAppResponse) SetStatusCode(v int32) *RecoverAppResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverAppResponse) SetBody(v *RecoverAppResponseBody) *RecoverAppResponse {
	s.Body = v
	return s
}

func (s *RecoverAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
