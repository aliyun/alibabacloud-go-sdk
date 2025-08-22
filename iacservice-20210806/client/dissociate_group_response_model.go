// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateGroupResponse
	GetStatusCode() *int32
	SetBody(v *DissociateGroupResponseBody) *DissociateGroupResponse
	GetBody() *DissociateGroupResponseBody
}

type DissociateGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateGroupResponse) GoString() string {
	return s.String()
}

func (s *DissociateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateGroupResponse) GetBody() *DissociateGroupResponseBody {
	return s.Body
}

func (s *DissociateGroupResponse) SetHeaders(v map[string]*string) *DissociateGroupResponse {
	s.Headers = v
	return s
}

func (s *DissociateGroupResponse) SetStatusCode(v int32) *DissociateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateGroupResponse) SetBody(v *DissociateGroupResponseBody) *DissociateGroupResponse {
	s.Body = v
	return s
}

func (s *DissociateGroupResponse) Validate() error {
	return dara.Validate(s)
}
