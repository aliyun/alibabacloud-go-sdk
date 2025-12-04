// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInfoForPodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeInfoForPodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeInfoForPodResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeInfoForPodResponseBody) *GetNodeInfoForPodResponse
	GetBody() *GetNodeInfoForPodResponseBody
}

type GetNodeInfoForPodResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeInfoForPodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeInfoForPodResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInfoForPodResponse) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeInfoForPodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeInfoForPodResponse) GetBody() *GetNodeInfoForPodResponseBody {
	return s.Body
}

func (s *GetNodeInfoForPodResponse) SetHeaders(v map[string]*string) *GetNodeInfoForPodResponse {
	s.Headers = v
	return s
}

func (s *GetNodeInfoForPodResponse) SetStatusCode(v int32) *GetNodeInfoForPodResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeInfoForPodResponse) SetBody(v *GetNodeInfoForPodResponseBody) *GetNodeInfoForPodResponse {
	s.Body = v
	return s
}

func (s *GetNodeInfoForPodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
