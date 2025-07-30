// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUpgradableVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUpgradableVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUpgradableVersionsResponse
	GetStatusCode() *int32
	SetBody(v *QueryUpgradableVersionsResponseBody) *QueryUpgradableVersionsResponse
	GetBody() *QueryUpgradableVersionsResponseBody
}

type QueryUpgradableVersionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUpgradableVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUpgradableVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUpgradableVersionsResponse) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUpgradableVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUpgradableVersionsResponse) GetBody() *QueryUpgradableVersionsResponseBody {
	return s.Body
}

func (s *QueryUpgradableVersionsResponse) SetHeaders(v map[string]*string) *QueryUpgradableVersionsResponse {
	s.Headers = v
	return s
}

func (s *QueryUpgradableVersionsResponse) SetStatusCode(v int32) *QueryUpgradableVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponse) SetBody(v *QueryUpgradableVersionsResponseBody) *QueryUpgradableVersionsResponse {
	s.Body = v
	return s
}

func (s *QueryUpgradableVersionsResponse) Validate() error {
	return dara.Validate(s)
}
