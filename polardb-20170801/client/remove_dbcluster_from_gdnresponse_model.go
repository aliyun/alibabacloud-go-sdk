// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDBClusterFromGDNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDBClusterFromGDNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDBClusterFromGDNResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDBClusterFromGDNResponseBody) *RemoveDBClusterFromGDNResponse
	GetBody() *RemoveDBClusterFromGDNResponseBody
}

type RemoveDBClusterFromGDNResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDBClusterFromGDNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDBClusterFromGDNResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDBClusterFromGDNResponse) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDBClusterFromGDNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDBClusterFromGDNResponse) GetBody() *RemoveDBClusterFromGDNResponseBody {
	return s.Body
}

func (s *RemoveDBClusterFromGDNResponse) SetHeaders(v map[string]*string) *RemoveDBClusterFromGDNResponse {
	s.Headers = v
	return s
}

func (s *RemoveDBClusterFromGDNResponse) SetStatusCode(v int32) *RemoveDBClusterFromGDNResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDBClusterFromGDNResponse) SetBody(v *RemoveDBClusterFromGDNResponseBody) *RemoveDBClusterFromGDNResponse {
	s.Body = v
	return s
}

func (s *RemoveDBClusterFromGDNResponse) Validate() error {
	return dara.Validate(s)
}
