// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineNamepaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEngineNamepaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEngineNamepaceResponse
	GetStatusCode() *int32
	SetBody(v *GetEngineNamepaceResponseBody) *GetEngineNamepaceResponse
	GetBody() *GetEngineNamepaceResponseBody
}

type GetEngineNamepaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEngineNamepaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEngineNamepaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEngineNamepaceResponse) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEngineNamepaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEngineNamepaceResponse) GetBody() *GetEngineNamepaceResponseBody {
	return s.Body
}

func (s *GetEngineNamepaceResponse) SetHeaders(v map[string]*string) *GetEngineNamepaceResponse {
	s.Headers = v
	return s
}

func (s *GetEngineNamepaceResponse) SetStatusCode(v int32) *GetEngineNamepaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEngineNamepaceResponse) SetBody(v *GetEngineNamepaceResponseBody) *GetEngineNamepaceResponse {
	s.Body = v
	return s
}

func (s *GetEngineNamepaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
