// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetStorageACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetStorageACLResponse
	GetStatusCode() *int32
	SetBody(v *SetStorageACLResponseBody) *SetStorageACLResponse
	GetBody() *SetStorageACLResponseBody
}

type SetStorageACLResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetStorageACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetStorageACLResponse) String() string {
	return dara.Prettify(s)
}

func (s SetStorageACLResponse) GoString() string {
	return s.String()
}

func (s *SetStorageACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetStorageACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetStorageACLResponse) GetBody() *SetStorageACLResponseBody {
	return s.Body
}

func (s *SetStorageACLResponse) SetHeaders(v map[string]*string) *SetStorageACLResponse {
	s.Headers = v
	return s
}

func (s *SetStorageACLResponse) SetStatusCode(v int32) *SetStorageACLResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStorageACLResponse) SetBody(v *SetStorageACLResponseBody) *SetStorageACLResponse {
	s.Body = v
	return s
}

func (s *SetStorageACLResponse) Validate() error {
	return dara.Validate(s)
}
