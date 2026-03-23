// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAddressByMacResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApAddressByMacResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApAddressByMacResponse
	GetStatusCode() *int32
	SetBody(v *GetApAddressByMacResponseBody) *GetApAddressByMacResponse
	GetBody() *GetApAddressByMacResponseBody
}

type GetApAddressByMacResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApAddressByMacResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApAddressByMacResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApAddressByMacResponse) GoString() string {
	return s.String()
}

func (s *GetApAddressByMacResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApAddressByMacResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApAddressByMacResponse) GetBody() *GetApAddressByMacResponseBody {
	return s.Body
}

func (s *GetApAddressByMacResponse) SetHeaders(v map[string]*string) *GetApAddressByMacResponse {
	s.Headers = v
	return s
}

func (s *GetApAddressByMacResponse) SetStatusCode(v int32) *GetApAddressByMacResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApAddressByMacResponse) SetBody(v *GetApAddressByMacResponseBody) *GetApAddressByMacResponse {
	s.Body = v
	return s
}

func (s *GetApAddressByMacResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
