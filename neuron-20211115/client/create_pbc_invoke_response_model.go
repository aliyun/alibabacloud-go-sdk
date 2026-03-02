// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInvokeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcInvokeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcInvokeResponse
	GetStatusCode() *int32
	SetBody(v *CreatePbcInvokeResponseBody) *CreatePbcInvokeResponse
	GetBody() *CreatePbcInvokeResponseBody
}

type CreatePbcInvokeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePbcInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcInvokeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInvokeResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcInvokeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcInvokeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcInvokeResponse) GetBody() *CreatePbcInvokeResponseBody {
	return s.Body
}

func (s *CreatePbcInvokeResponse) SetHeaders(v map[string]*string) *CreatePbcInvokeResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcInvokeResponse) SetStatusCode(v int32) *CreatePbcInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcInvokeResponse) SetBody(v *CreatePbcInvokeResponseBody) *CreatePbcInvokeResponse {
	s.Body = v
	return s
}

func (s *CreatePbcInvokeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
