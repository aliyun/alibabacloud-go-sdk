// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormDataByIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFormDataByIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFormDataByIDResponse
	GetStatusCode() *int32
	SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse
	GetBody() *GetFormDataByIDResponseBody
}

type GetFormDataByIDResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormDataByIDResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDResponse) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFormDataByIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFormDataByIDResponse) GetBody() *GetFormDataByIDResponseBody {
	return s.Body
}

func (s *GetFormDataByIDResponse) SetHeaders(v map[string]*string) *GetFormDataByIDResponse {
	s.Headers = v
	return s
}

func (s *GetFormDataByIDResponse) SetStatusCode(v int32) *GetFormDataByIDResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
	s.Body = v
	return s
}

func (s *GetFormDataByIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
