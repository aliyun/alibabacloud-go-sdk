// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupAddOrUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUserGroupAddOrUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUserGroupAddOrUpdateResponse
	GetStatusCode() *int32
	SetBody(v *DsgUserGroupAddOrUpdateResponseBody) *DsgUserGroupAddOrUpdateResponse
	GetBody() *DsgUserGroupAddOrUpdateResponseBody
}

type DsgUserGroupAddOrUpdateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUserGroupAddOrUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUserGroupAddOrUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupAddOrUpdateResponse) GoString() string {
	return s.String()
}

func (s *DsgUserGroupAddOrUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUserGroupAddOrUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUserGroupAddOrUpdateResponse) GetBody() *DsgUserGroupAddOrUpdateResponseBody {
	return s.Body
}

func (s *DsgUserGroupAddOrUpdateResponse) SetHeaders(v map[string]*string) *DsgUserGroupAddOrUpdateResponse {
	s.Headers = v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponse) SetStatusCode(v int32) *DsgUserGroupAddOrUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponse) SetBody(v *DsgUserGroupAddOrUpdateResponseBody) *DsgUserGroupAddOrUpdateResponse {
	s.Body = v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponse) Validate() error {
	return dara.Validate(s)
}
