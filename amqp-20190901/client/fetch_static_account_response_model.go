// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchStaticAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchStaticAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchStaticAccountResponse
	GetStatusCode() *int32
	SetBody(v *FetchStaticAccountResponseBody) *FetchStaticAccountResponse
	GetBody() *FetchStaticAccountResponseBody
}

type FetchStaticAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchStaticAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchStaticAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchStaticAccountResponse) GoString() string {
	return s.String()
}

func (s *FetchStaticAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchStaticAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchStaticAccountResponse) GetBody() *FetchStaticAccountResponseBody {
	return s.Body
}

func (s *FetchStaticAccountResponse) SetHeaders(v map[string]*string) *FetchStaticAccountResponse {
	s.Headers = v
	return s
}

func (s *FetchStaticAccountResponse) SetStatusCode(v int32) *FetchStaticAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchStaticAccountResponse) SetBody(v *FetchStaticAccountResponseBody) *FetchStaticAccountResponse {
	s.Body = v
	return s
}

func (s *FetchStaticAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
