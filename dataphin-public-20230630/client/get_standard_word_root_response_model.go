// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardWordRootResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardWordRootResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardWordRootResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardWordRootResponseBody) *GetStandardWordRootResponse
	GetBody() *GetStandardWordRootResponseBody
}

type GetStandardWordRootResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardWordRootResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardWordRootResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardWordRootResponse) GoString() string {
	return s.String()
}

func (s *GetStandardWordRootResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardWordRootResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardWordRootResponse) GetBody() *GetStandardWordRootResponseBody {
	return s.Body
}

func (s *GetStandardWordRootResponse) SetHeaders(v map[string]*string) *GetStandardWordRootResponse {
	s.Headers = v
	return s
}

func (s *GetStandardWordRootResponse) SetStatusCode(v int32) *GetStandardWordRootResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardWordRootResponse) SetBody(v *GetStandardWordRootResponseBody) *GetStandardWordRootResponse {
	s.Body = v
	return s
}

func (s *GetStandardWordRootResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
