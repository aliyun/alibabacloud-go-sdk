// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectableClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddConnectableClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddConnectableClusterResponse
	GetStatusCode() *int32
	SetBody(v *AddConnectableClusterResponseBody) *AddConnectableClusterResponse
	GetBody() *AddConnectableClusterResponseBody
}

type AddConnectableClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddConnectableClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddConnectableClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s AddConnectableClusterResponse) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddConnectableClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddConnectableClusterResponse) GetBody() *AddConnectableClusterResponseBody {
	return s.Body
}

func (s *AddConnectableClusterResponse) SetHeaders(v map[string]*string) *AddConnectableClusterResponse {
	s.Headers = v
	return s
}

func (s *AddConnectableClusterResponse) SetStatusCode(v int32) *AddConnectableClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddConnectableClusterResponse) SetBody(v *AddConnectableClusterResponseBody) *AddConnectableClusterResponse {
	s.Body = v
	return s
}

func (s *AddConnectableClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
