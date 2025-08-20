// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackGroupOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackGroupOperationResponse
	GetStatusCode() *int32
	SetBody(v *GetStackGroupOperationResponseBody) *GetStackGroupOperationResponse
	GetBody() *GetStackGroupOperationResponseBody
}

type GetStackGroupOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackGroupOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackGroupOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackGroupOperationResponse) GetBody() *GetStackGroupOperationResponseBody {
	return s.Body
}

func (s *GetStackGroupOperationResponse) SetHeaders(v map[string]*string) *GetStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupOperationResponse) SetStatusCode(v int32) *GetStackGroupOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackGroupOperationResponse) SetBody(v *GetStackGroupOperationResponseBody) *GetStackGroupOperationResponse {
	s.Body = v
	return s
}

func (s *GetStackGroupOperationResponse) Validate() error {
	return dara.Validate(s)
}
