// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServiceClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnsServiceClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnsServiceClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListAnsServiceClustersResponseBody) *ListAnsServiceClustersResponse
	GetBody() *ListAnsServiceClustersResponseBody
}

type ListAnsServiceClustersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnsServiceClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnsServiceClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersResponse) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnsServiceClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnsServiceClustersResponse) GetBody() *ListAnsServiceClustersResponseBody {
	return s.Body
}

func (s *ListAnsServiceClustersResponse) SetHeaders(v map[string]*string) *ListAnsServiceClustersResponse {
	s.Headers = v
	return s
}

func (s *ListAnsServiceClustersResponse) SetStatusCode(v int32) *ListAnsServiceClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnsServiceClustersResponse) SetBody(v *ListAnsServiceClustersResponseBody) *ListAnsServiceClustersResponse {
	s.Body = v
	return s
}

func (s *ListAnsServiceClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
