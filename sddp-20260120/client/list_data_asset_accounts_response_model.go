// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAssetAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAssetAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAssetAccountsResponseBody) *ListDataAssetAccountsResponse
	GetBody() *ListDataAssetAccountsResponseBody
}

type ListDataAssetAccountsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAssetAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAssetAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListDataAssetAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAssetAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAssetAccountsResponse) GetBody() *ListDataAssetAccountsResponseBody {
	return s.Body
}

func (s *ListDataAssetAccountsResponse) SetHeaders(v map[string]*string) *ListDataAssetAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListDataAssetAccountsResponse) SetStatusCode(v int32) *ListDataAssetAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAssetAccountsResponse) SetBody(v *ListDataAssetAccountsResponseBody) *ListDataAssetAccountsResponse {
	s.Body = v
	return s
}

func (s *ListDataAssetAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
