// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingErrorByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBindingErrorByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBindingErrorByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetBindingErrorByTaskIdResponseBody) *GetBindingErrorByTaskIdResponse
	GetBody() *GetBindingErrorByTaskIdResponseBody
}

type GetBindingErrorByTaskIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBindingErrorByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBindingErrorByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBindingErrorByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBindingErrorByTaskIdResponse) GetBody() *GetBindingErrorByTaskIdResponseBody {
	return s.Body
}

func (s *GetBindingErrorByTaskIdResponse) SetHeaders(v map[string]*string) *GetBindingErrorByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetBindingErrorByTaskIdResponse) SetStatusCode(v int32) *GetBindingErrorByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponse) SetBody(v *GetBindingErrorByTaskIdResponseBody) *GetBindingErrorByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetBindingErrorByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
