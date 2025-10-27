// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageEventOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageEventOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageEventOperationResponse
	GetStatusCode() *int32
	SetBody(v *AddImageEventOperationResponseBody) *AddImageEventOperationResponse
	GetBody() *AddImageEventOperationResponseBody
}

type AddImageEventOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageEventOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageEventOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageEventOperationResponse) GoString() string {
	return s.String()
}

func (s *AddImageEventOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageEventOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageEventOperationResponse) GetBody() *AddImageEventOperationResponseBody {
	return s.Body
}

func (s *AddImageEventOperationResponse) SetHeaders(v map[string]*string) *AddImageEventOperationResponse {
	s.Headers = v
	return s
}

func (s *AddImageEventOperationResponse) SetStatusCode(v int32) *AddImageEventOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageEventOperationResponse) SetBody(v *AddImageEventOperationResponseBody) *AddImageEventOperationResponse {
	s.Body = v
	return s
}

func (s *AddImageEventOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
