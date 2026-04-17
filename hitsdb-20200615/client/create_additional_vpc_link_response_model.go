// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdditionalVpcLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdditionalVpcLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdditionalVpcLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdditionalVpcLinkResponseBody) *CreateAdditionalVpcLinkResponse
	GetBody() *CreateAdditionalVpcLinkResponseBody
}

type CreateAdditionalVpcLinkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdditionalVpcLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdditionalVpcLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdditionalVpcLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateAdditionalVpcLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdditionalVpcLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdditionalVpcLinkResponse) GetBody() *CreateAdditionalVpcLinkResponseBody {
	return s.Body
}

func (s *CreateAdditionalVpcLinkResponse) SetHeaders(v map[string]*string) *CreateAdditionalVpcLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateAdditionalVpcLinkResponse) SetStatusCode(v int32) *CreateAdditionalVpcLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdditionalVpcLinkResponse) SetBody(v *CreateAdditionalVpcLinkResponseBody) *CreateAdditionalVpcLinkResponse {
	s.Body = v
	return s
}

func (s *CreateAdditionalVpcLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
