// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcPrefixListEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcPrefixListEntriesResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcPrefixListEntriesResponseBody) *GetVpcPrefixListEntriesResponse
	GetBody() *GetVpcPrefixListEntriesResponseBody
}

type GetVpcPrefixListEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcPrefixListEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcPrefixListEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListEntriesResponse) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcPrefixListEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcPrefixListEntriesResponse) GetBody() *GetVpcPrefixListEntriesResponseBody {
	return s.Body
}

func (s *GetVpcPrefixListEntriesResponse) SetHeaders(v map[string]*string) *GetVpcPrefixListEntriesResponse {
	s.Headers = v
	return s
}

func (s *GetVpcPrefixListEntriesResponse) SetStatusCode(v int32) *GetVpcPrefixListEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponse) SetBody(v *GetVpcPrefixListEntriesResponseBody) *GetVpcPrefixListEntriesResponse {
	s.Body = v
	return s
}

func (s *GetVpcPrefixListEntriesResponse) Validate() error {
	return dara.Validate(s)
}
