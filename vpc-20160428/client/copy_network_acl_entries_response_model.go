// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyNetworkAclEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyNetworkAclEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyNetworkAclEntriesResponse
	GetStatusCode() *int32
	SetBody(v *CopyNetworkAclEntriesResponseBody) *CopyNetworkAclEntriesResponse
	GetBody() *CopyNetworkAclEntriesResponseBody
}

type CopyNetworkAclEntriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyNetworkAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyNetworkAclEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyNetworkAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *CopyNetworkAclEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyNetworkAclEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyNetworkAclEntriesResponse) GetBody() *CopyNetworkAclEntriesResponseBody {
	return s.Body
}

func (s *CopyNetworkAclEntriesResponse) SetHeaders(v map[string]*string) *CopyNetworkAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *CopyNetworkAclEntriesResponse) SetStatusCode(v int32) *CopyNetworkAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyNetworkAclEntriesResponse) SetBody(v *CopyNetworkAclEntriesResponseBody) *CopyNetworkAclEntriesResponse {
	s.Body = v
	return s
}

func (s *CopyNetworkAclEntriesResponse) Validate() error {
	return dara.Validate(s)
}
