// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMemberAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMemberAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMemberAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMemberAttributesResponseBody) *ModifyInstanceMemberAttributesResponse
	GetBody() *ModifyInstanceMemberAttributesResponseBody
}

type ModifyInstanceMemberAttributesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMemberAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMemberAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMemberAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMemberAttributesResponse) GetBody() *ModifyInstanceMemberAttributesResponseBody {
	return s.Body
}

func (s *ModifyInstanceMemberAttributesResponse) SetHeaders(v map[string]*string) *ModifyInstanceMemberAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetStatusCode(v int32) *ModifyInstanceMemberAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetBody(v *ModifyInstanceMemberAttributesResponseBody) *ModifyInstanceMemberAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
