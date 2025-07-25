// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainGroupResponseBody) *DeleteDomainGroupResponse
	GetBody() *DeleteDomainGroupResponseBody
}

type DeleteDomainGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainGroupResponse) GetBody() *DeleteDomainGroupResponseBody {
	return s.Body
}

func (s *DeleteDomainGroupResponse) SetHeaders(v map[string]*string) *DeleteDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainGroupResponse) SetStatusCode(v int32) *DeleteDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainGroupResponse) SetBody(v *DeleteDomainGroupResponseBody) *DeleteDomainGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainGroupResponse) Validate() error {
	return dara.Validate(s)
}
