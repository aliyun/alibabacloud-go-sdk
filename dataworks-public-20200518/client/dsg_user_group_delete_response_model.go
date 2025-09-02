// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUserGroupDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUserGroupDeleteResponse
	GetStatusCode() *int32
	SetBody(v *DsgUserGroupDeleteResponseBody) *DsgUserGroupDeleteResponse
	GetBody() *DsgUserGroupDeleteResponseBody
}

type DsgUserGroupDeleteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUserGroupDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUserGroupDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupDeleteResponse) GoString() string {
	return s.String()
}

func (s *DsgUserGroupDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUserGroupDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUserGroupDeleteResponse) GetBody() *DsgUserGroupDeleteResponseBody {
	return s.Body
}

func (s *DsgUserGroupDeleteResponse) SetHeaders(v map[string]*string) *DsgUserGroupDeleteResponse {
	s.Headers = v
	return s
}

func (s *DsgUserGroupDeleteResponse) SetStatusCode(v int32) *DsgUserGroupDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUserGroupDeleteResponse) SetBody(v *DsgUserGroupDeleteResponseBody) *DsgUserGroupDeleteResponse {
	s.Body = v
	return s
}

func (s *DsgUserGroupDeleteResponse) Validate() error {
	return dara.Validate(s)
}
