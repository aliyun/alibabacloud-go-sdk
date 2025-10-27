// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContainerNetworkConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindContainerNetworkConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindContainerNetworkConnectResponse
	GetStatusCode() *int32
	SetBody(v *FindContainerNetworkConnectResponseBody) *FindContainerNetworkConnectResponse
	GetBody() *FindContainerNetworkConnectResponseBody
}

type FindContainerNetworkConnectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindContainerNetworkConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindContainerNetworkConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectResponse) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindContainerNetworkConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindContainerNetworkConnectResponse) GetBody() *FindContainerNetworkConnectResponseBody {
	return s.Body
}

func (s *FindContainerNetworkConnectResponse) SetHeaders(v map[string]*string) *FindContainerNetworkConnectResponse {
	s.Headers = v
	return s
}

func (s *FindContainerNetworkConnectResponse) SetStatusCode(v int32) *FindContainerNetworkConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *FindContainerNetworkConnectResponse) SetBody(v *FindContainerNetworkConnectResponseBody) *FindContainerNetworkConnectResponse {
	s.Body = v
	return s
}

func (s *FindContainerNetworkConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
