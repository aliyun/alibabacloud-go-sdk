// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcPrefixListAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcPrefixListAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcPrefixListAssociationsResponseBody) *GetVpcPrefixListAssociationsResponse
	GetBody() *GetVpcPrefixListAssociationsResponseBody
}

type GetVpcPrefixListAssociationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcPrefixListAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcPrefixListAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListAssociationsResponse) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcPrefixListAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcPrefixListAssociationsResponse) GetBody() *GetVpcPrefixListAssociationsResponseBody {
	return s.Body
}

func (s *GetVpcPrefixListAssociationsResponse) SetHeaders(v map[string]*string) *GetVpcPrefixListAssociationsResponse {
	s.Headers = v
	return s
}

func (s *GetVpcPrefixListAssociationsResponse) SetStatusCode(v int32) *GetVpcPrefixListAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponse) SetBody(v *GetVpcPrefixListAssociationsResponseBody) *GetVpcPrefixListAssociationsResponse {
	s.Body = v
	return s
}

func (s *GetVpcPrefixListAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
