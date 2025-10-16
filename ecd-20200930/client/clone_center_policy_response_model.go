// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneCenterPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneCenterPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneCenterPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CloneCenterPolicyResponseBody) *CloneCenterPolicyResponse
	GetBody() *CloneCenterPolicyResponseBody
}

type CloneCenterPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneCenterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneCenterPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneCenterPolicyResponse) GoString() string {
	return s.String()
}

func (s *CloneCenterPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneCenterPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneCenterPolicyResponse) GetBody() *CloneCenterPolicyResponseBody {
	return s.Body
}

func (s *CloneCenterPolicyResponse) SetHeaders(v map[string]*string) *CloneCenterPolicyResponse {
	s.Headers = v
	return s
}

func (s *CloneCenterPolicyResponse) SetStatusCode(v int32) *CloneCenterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneCenterPolicyResponse) SetBody(v *CloneCenterPolicyResponseBody) *CloneCenterPolicyResponse {
	s.Body = v
	return s
}

func (s *CloneCenterPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
