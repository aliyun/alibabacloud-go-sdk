// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAckClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAckClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListAckClustersResponseBody) *ListAckClustersResponse
	GetBody() *ListAckClustersResponseBody
}

type ListAckClustersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAckClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAckClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAckClustersResponse) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAckClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAckClustersResponse) GetBody() *ListAckClustersResponseBody {
	return s.Body
}

func (s *ListAckClustersResponse) SetHeaders(v map[string]*string) *ListAckClustersResponse {
	s.Headers = v
	return s
}

func (s *ListAckClustersResponse) SetStatusCode(v int32) *ListAckClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAckClustersResponse) SetBody(v *ListAckClustersResponseBody) *ListAckClustersResponse {
	s.Body = v
	return s
}

func (s *ListAckClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
