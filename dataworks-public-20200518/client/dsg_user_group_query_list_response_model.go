// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupQueryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUserGroupQueryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUserGroupQueryListResponse
	GetStatusCode() *int32
	SetBody(v *DsgUserGroupQueryListResponseBody) *DsgUserGroupQueryListResponse
	GetBody() *DsgUserGroupQueryListResponseBody
}

type DsgUserGroupQueryListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUserGroupQueryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUserGroupQueryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryListResponse) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUserGroupQueryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUserGroupQueryListResponse) GetBody() *DsgUserGroupQueryListResponseBody {
	return s.Body
}

func (s *DsgUserGroupQueryListResponse) SetHeaders(v map[string]*string) *DsgUserGroupQueryListResponse {
	s.Headers = v
	return s
}

func (s *DsgUserGroupQueryListResponse) SetStatusCode(v int32) *DsgUserGroupQueryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUserGroupQueryListResponse) SetBody(v *DsgUserGroupQueryListResponseBody) *DsgUserGroupQueryListResponse {
	s.Body = v
	return s
}

func (s *DsgUserGroupQueryListResponse) Validate() error {
	return dara.Validate(s)
}
