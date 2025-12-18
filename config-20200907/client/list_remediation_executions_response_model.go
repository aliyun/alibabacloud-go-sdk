// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemediationExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemediationExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRemediationExecutionsResponseBody) *ListRemediationExecutionsResponse
	GetBody() *ListRemediationExecutionsResponseBody
}

type ListRemediationExecutionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemediationExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemediationExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListRemediationExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemediationExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemediationExecutionsResponse) GetBody() *ListRemediationExecutionsResponseBody {
	return s.Body
}

func (s *ListRemediationExecutionsResponse) SetHeaders(v map[string]*string) *ListRemediationExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListRemediationExecutionsResponse) SetStatusCode(v int32) *ListRemediationExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemediationExecutionsResponse) SetBody(v *ListRemediationExecutionsResponseBody) *ListRemediationExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListRemediationExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
