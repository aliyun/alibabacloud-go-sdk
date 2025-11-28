// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSystemContractResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllSystemContractResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllSystemContractResponse
	GetStatusCode() *int32
	SetBody(v *ListAllSystemContractResponseBody) *ListAllSystemContractResponse
	GetBody() *ListAllSystemContractResponseBody
}

type ListAllSystemContractResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllSystemContractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllSystemContractResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllSystemContractResponse) GoString() string {
	return s.String()
}

func (s *ListAllSystemContractResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllSystemContractResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllSystemContractResponse) GetBody() *ListAllSystemContractResponseBody {
	return s.Body
}

func (s *ListAllSystemContractResponse) SetHeaders(v map[string]*string) *ListAllSystemContractResponse {
	s.Headers = v
	return s
}

func (s *ListAllSystemContractResponse) SetStatusCode(v int32) *ListAllSystemContractResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllSystemContractResponse) SetBody(v *ListAllSystemContractResponseBody) *ListAllSystemContractResponse {
	s.Body = v
	return s
}

func (s *ListAllSystemContractResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
