// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterTransferToMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterTransferToMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterTransferToMemberResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterTransferToMemberResponseBody) *ModelRouterTransferToMemberResponse
	GetBody() *ModelRouterTransferToMemberResponseBody
}

type ModelRouterTransferToMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterTransferToMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterTransferToMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterTransferToMemberResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterTransferToMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterTransferToMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterTransferToMemberResponse) GetBody() *ModelRouterTransferToMemberResponseBody {
	return s.Body
}

func (s *ModelRouterTransferToMemberResponse) SetHeaders(v map[string]*string) *ModelRouterTransferToMemberResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterTransferToMemberResponse) SetStatusCode(v int32) *ModelRouterTransferToMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterTransferToMemberResponse) SetBody(v *ModelRouterTransferToMemberResponseBody) *ModelRouterTransferToMemberResponse {
	s.Body = v
	return s
}

func (s *ModelRouterTransferToMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
