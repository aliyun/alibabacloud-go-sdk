// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetElecConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetElecConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetElecConstituteResponseBody) *GetElecConstituteResponse
	GetBody() *GetElecConstituteResponseBody
}

type GetElecConstituteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElecConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElecConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetElecConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetElecConstituteResponse) GetBody() *GetElecConstituteResponseBody {
	return s.Body
}

func (s *GetElecConstituteResponse) SetHeaders(v map[string]*string) *GetElecConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetElecConstituteResponse) SetStatusCode(v int32) *GetElecConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElecConstituteResponse) SetBody(v *GetElecConstituteResponseBody) *GetElecConstituteResponse {
	s.Body = v
	return s
}

func (s *GetElecConstituteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
