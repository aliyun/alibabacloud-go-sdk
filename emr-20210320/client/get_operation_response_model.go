// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationResponseBody) *GetOperationResponse
	GetBody() *GetOperationResponseBody
}

type GetOperationResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationResponse) GoString() string {
	return s.String()
}

func (s *GetOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationResponse) GetBody() *GetOperationResponseBody {
	return s.Body
}

func (s *GetOperationResponse) SetHeaders(v map[string]*string) *GetOperationResponse {
	s.Headers = v
	return s
}

func (s *GetOperationResponse) SetStatusCode(v int32) *GetOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationResponse) SetBody(v *GetOperationResponseBody) *GetOperationResponse {
	s.Body = v
	return s
}

func (s *GetOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
