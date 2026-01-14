// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicAcceleratorResponseBody) *GetBasicAcceleratorResponse
	GetBody() *GetBasicAcceleratorResponseBody
}

type GetBasicAcceleratorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicAcceleratorResponse) GetBody() *GetBasicAcceleratorResponseBody {
	return s.Body
}

func (s *GetBasicAcceleratorResponse) SetHeaders(v map[string]*string) *GetBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAcceleratorResponse) SetStatusCode(v int32) *GetBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAcceleratorResponse) SetBody(v *GetBasicAcceleratorResponseBody) *GetBasicAcceleratorResponse {
	s.Body = v
	return s
}

func (s *GetBasicAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
