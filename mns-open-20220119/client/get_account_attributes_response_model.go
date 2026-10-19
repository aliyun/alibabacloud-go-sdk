// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountAttributesResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountAttributesResponseBody) *GetAccountAttributesResponse
	GetBody() *GetAccountAttributesResponseBody
}

type GetAccountAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetAccountAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountAttributesResponse) GetBody() *GetAccountAttributesResponseBody {
	return s.Body
}

func (s *GetAccountAttributesResponse) SetHeaders(v map[string]*string) *GetAccountAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetAccountAttributesResponse) SetStatusCode(v int32) *GetAccountAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountAttributesResponse) SetBody(v *GetAccountAttributesResponseBody) *GetAccountAttributesResponse {
	s.Body = v
	return s
}

func (s *GetAccountAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
