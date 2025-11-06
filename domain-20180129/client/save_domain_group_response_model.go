// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *SaveDomainGroupResponseBody) *SaveDomainGroupResponse
	GetBody() *SaveDomainGroupResponseBody
}

type SaveDomainGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveDomainGroupResponse) GetBody() *SaveDomainGroupResponseBody {
	return s.Body
}

func (s *SaveDomainGroupResponse) SetHeaders(v map[string]*string) *SaveDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *SaveDomainGroupResponse) SetStatusCode(v int32) *SaveDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDomainGroupResponse) SetBody(v *SaveDomainGroupResponseBody) *SaveDomainGroupResponse {
	s.Body = v
	return s
}

func (s *SaveDomainGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
