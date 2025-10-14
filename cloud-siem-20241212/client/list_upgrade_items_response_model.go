// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUpgradeItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUpgradeItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListUpgradeItemsResponseBody) *ListUpgradeItemsResponse
	GetBody() *ListUpgradeItemsResponseBody
}

type ListUpgradeItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUpgradeItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUpgradeItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeItemsResponse) GoString() string {
	return s.String()
}

func (s *ListUpgradeItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUpgradeItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUpgradeItemsResponse) GetBody() *ListUpgradeItemsResponseBody {
	return s.Body
}

func (s *ListUpgradeItemsResponse) SetHeaders(v map[string]*string) *ListUpgradeItemsResponse {
	s.Headers = v
	return s
}

func (s *ListUpgradeItemsResponse) SetStatusCode(v int32) *ListUpgradeItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpgradeItemsResponse) SetBody(v *ListUpgradeItemsResponseBody) *ListUpgradeItemsResponse {
	s.Body = v
	return s
}

func (s *ListUpgradeItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
