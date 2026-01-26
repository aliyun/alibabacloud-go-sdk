// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommercialStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommercialStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommercialStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetCommercialStatusResponseBody) *GetCommercialStatusResponse
	GetBody() *GetCommercialStatusResponseBody
}

type GetCommercialStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommercialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommercialStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommercialStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCommercialStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommercialStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommercialStatusResponse) GetBody() *GetCommercialStatusResponseBody {
	return s.Body
}

func (s *GetCommercialStatusResponse) SetHeaders(v map[string]*string) *GetCommercialStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCommercialStatusResponse) SetStatusCode(v int32) *GetCommercialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommercialStatusResponse) SetBody(v *GetCommercialStatusResponseBody) *GetCommercialStatusResponse {
	s.Body = v
	return s
}

func (s *GetCommercialStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
