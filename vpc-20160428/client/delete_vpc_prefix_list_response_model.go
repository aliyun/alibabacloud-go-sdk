// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcPrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcPrefixListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcPrefixListResponseBody) *DeleteVpcPrefixListResponse
	GetBody() *DeleteVpcPrefixListResponseBody
}

type DeleteVpcPrefixListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcPrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcPrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPrefixListResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcPrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcPrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcPrefixListResponse) GetBody() *DeleteVpcPrefixListResponseBody {
	return s.Body
}

func (s *DeleteVpcPrefixListResponse) SetHeaders(v map[string]*string) *DeleteVpcPrefixListResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcPrefixListResponse) SetStatusCode(v int32) *DeleteVpcPrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcPrefixListResponse) SetBody(v *DeleteVpcPrefixListResponseBody) *DeleteVpcPrefixListResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcPrefixListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
