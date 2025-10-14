// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDomainGroupResponseBody) *ChangeDomainGroupResponse
	GetBody() *ChangeDomainGroupResponseBody
}

type ChangeDomainGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDomainGroupResponse) GetBody() *ChangeDomainGroupResponseBody {
	return s.Body
}

func (s *ChangeDomainGroupResponse) SetHeaders(v map[string]*string) *ChangeDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeDomainGroupResponse) SetStatusCode(v int32) *ChangeDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDomainGroupResponse) SetBody(v *ChangeDomainGroupResponseBody) *ChangeDomainGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeDomainGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
