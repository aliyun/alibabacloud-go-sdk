// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRCVClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRCVClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRCVClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListRCVClustersResponseBody) *ListRCVClustersResponse
	GetBody() *ListRCVClustersResponseBody
}

type ListRCVClustersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRCVClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRCVClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponse) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRCVClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRCVClustersResponse) GetBody() *ListRCVClustersResponseBody {
	return s.Body
}

func (s *ListRCVClustersResponse) SetHeaders(v map[string]*string) *ListRCVClustersResponse {
	s.Headers = v
	return s
}

func (s *ListRCVClustersResponse) SetStatusCode(v int32) *ListRCVClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRCVClustersResponse) SetBody(v *ListRCVClustersResponseBody) *ListRCVClustersResponse {
	s.Body = v
	return s
}

func (s *ListRCVClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
