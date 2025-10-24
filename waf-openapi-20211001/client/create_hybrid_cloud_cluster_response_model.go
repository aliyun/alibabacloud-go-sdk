// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridCloudClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridCloudClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridCloudClusterResponseBody) *CreateHybridCloudClusterResponse
	GetBody() *CreateHybridCloudClusterResponseBody
}

type CreateHybridCloudClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridCloudClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridCloudClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridCloudClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridCloudClusterResponse) GetBody() *CreateHybridCloudClusterResponseBody {
	return s.Body
}

func (s *CreateHybridCloudClusterResponse) SetHeaders(v map[string]*string) *CreateHybridCloudClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridCloudClusterResponse) SetStatusCode(v int32) *CreateHybridCloudClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridCloudClusterResponse) SetBody(v *CreateHybridCloudClusterResponseBody) *CreateHybridCloudClusterResponse {
	s.Body = v
	return s
}

func (s *CreateHybridCloudClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
