// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBdrcServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBdrcServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBdrcServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetBdrcServiceResponseBody) *GetBdrcServiceResponse
	GetBody() *GetBdrcServiceResponseBody
}

type GetBdrcServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBdrcServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBdrcServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBdrcServiceResponse) GetBody() *GetBdrcServiceResponseBody {
	return s.Body
}

func (s *GetBdrcServiceResponse) SetHeaders(v map[string]*string) *GetBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *GetBdrcServiceResponse) SetStatusCode(v int32) *GetBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBdrcServiceResponse) SetBody(v *GetBdrcServiceResponseBody) *GetBdrcServiceResponse {
	s.Body = v
	return s
}

func (s *GetBdrcServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
