// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAccountAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearAccountAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearAccountAliasResponse
	GetStatusCode() *int32
	SetBody(v *ClearAccountAliasResponseBody) *ClearAccountAliasResponse
	GetBody() *ClearAccountAliasResponseBody
}

type ClearAccountAliasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAccountAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *ClearAccountAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearAccountAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearAccountAliasResponse) GetBody() *ClearAccountAliasResponseBody {
	return s.Body
}

func (s *ClearAccountAliasResponse) SetHeaders(v map[string]*string) *ClearAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *ClearAccountAliasResponse) SetStatusCode(v int32) *ClearAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAccountAliasResponse) SetBody(v *ClearAccountAliasResponseBody) *ClearAccountAliasResponse {
	s.Body = v
	return s
}

func (s *ClearAccountAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
