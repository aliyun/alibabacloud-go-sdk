// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSharedAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSharedAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSharedAccountsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSharedAccountsResponseBody) *RemoveSharedAccountsResponse
	GetBody() *RemoveSharedAccountsResponseBody
}

type RemoveSharedAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSharedAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *RemoveSharedAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSharedAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSharedAccountsResponse) GetBody() *RemoveSharedAccountsResponseBody {
	return s.Body
}

func (s *RemoveSharedAccountsResponse) SetHeaders(v map[string]*string) *RemoveSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *RemoveSharedAccountsResponse) SetStatusCode(v int32) *RemoveSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSharedAccountsResponse) SetBody(v *RemoveSharedAccountsResponseBody) *RemoveSharedAccountsResponse {
	s.Body = v
	return s
}

func (s *RemoveSharedAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
