// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSingleConnDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSingleConnDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSingleConnDataResponse
	GetStatusCode() *int32
	SetBody(v *GetSingleConnDataResponseBody) *GetSingleConnDataResponse
	GetBody() *GetSingleConnDataResponseBody
}

type GetSingleConnDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSingleConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSingleConnDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSingleConnDataResponse) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSingleConnDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSingleConnDataResponse) GetBody() *GetSingleConnDataResponseBody {
	return s.Body
}

func (s *GetSingleConnDataResponse) SetHeaders(v map[string]*string) *GetSingleConnDataResponse {
	s.Headers = v
	return s
}

func (s *GetSingleConnDataResponse) SetStatusCode(v int32) *GetSingleConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSingleConnDataResponse) SetBody(v *GetSingleConnDataResponseBody) *GetSingleConnDataResponse {
	s.Body = v
	return s
}

func (s *GetSingleConnDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
