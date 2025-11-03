// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcPrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcPrefixListResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcPrefixListResponseBody) *CreateVpcPrefixListResponse
	GetBody() *CreateVpcPrefixListResponseBody
}

type CreateVpcPrefixListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcPrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcPrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPrefixListResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcPrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcPrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcPrefixListResponse) GetBody() *CreateVpcPrefixListResponseBody {
	return s.Body
}

func (s *CreateVpcPrefixListResponse) SetHeaders(v map[string]*string) *CreateVpcPrefixListResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcPrefixListResponse) SetStatusCode(v int32) *CreateVpcPrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcPrefixListResponse) SetBody(v *CreateVpcPrefixListResponseBody) *CreateVpcPrefixListResponse {
	s.Body = v
	return s
}

func (s *CreateVpcPrefixListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
