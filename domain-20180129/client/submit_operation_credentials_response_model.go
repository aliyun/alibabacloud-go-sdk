// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitOperationCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitOperationCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitOperationCredentialsResponseBody) *SubmitOperationCredentialsResponse
	GetBody() *SubmitOperationCredentialsResponseBody
}

type SubmitOperationCredentialsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitOperationCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitOperationCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationCredentialsResponse) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitOperationCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitOperationCredentialsResponse) GetBody() *SubmitOperationCredentialsResponseBody {
	return s.Body
}

func (s *SubmitOperationCredentialsResponse) SetHeaders(v map[string]*string) *SubmitOperationCredentialsResponse {
	s.Headers = v
	return s
}

func (s *SubmitOperationCredentialsResponse) SetStatusCode(v int32) *SubmitOperationCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOperationCredentialsResponse) SetBody(v *SubmitOperationCredentialsResponseBody) *SubmitOperationCredentialsResponse {
	s.Body = v
	return s
}

func (s *SubmitOperationCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
