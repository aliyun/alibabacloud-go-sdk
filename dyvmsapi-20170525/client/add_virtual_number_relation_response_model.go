// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirtualNumberRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVirtualNumberRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVirtualNumberRelationResponse
	GetStatusCode() *int32
	SetBody(v *AddVirtualNumberRelationResponseBody) *AddVirtualNumberRelationResponse
	GetBody() *AddVirtualNumberRelationResponseBody
}

type AddVirtualNumberRelationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVirtualNumberRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVirtualNumberRelationResponse) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVirtualNumberRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVirtualNumberRelationResponse) GetBody() *AddVirtualNumberRelationResponseBody {
	return s.Body
}

func (s *AddVirtualNumberRelationResponse) SetHeaders(v map[string]*string) *AddVirtualNumberRelationResponse {
	s.Headers = v
	return s
}

func (s *AddVirtualNumberRelationResponse) SetStatusCode(v int32) *AddVirtualNumberRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVirtualNumberRelationResponse) SetBody(v *AddVirtualNumberRelationResponseBody) *AddVirtualNumberRelationResponse {
	s.Body = v
	return s
}

func (s *AddVirtualNumberRelationResponse) Validate() error {
	return dara.Validate(s)
}
