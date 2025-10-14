// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarxDataNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarxDataNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarxDataNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarxDataNodesResponseBody) *DescribePolarxDataNodesResponse
	GetBody() *DescribePolarxDataNodesResponseBody
}

type DescribePolarxDataNodesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarxDataNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarxDataNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarxDataNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarxDataNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarxDataNodesResponse) GetBody() *DescribePolarxDataNodesResponseBody {
	return s.Body
}

func (s *DescribePolarxDataNodesResponse) SetHeaders(v map[string]*string) *DescribePolarxDataNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarxDataNodesResponse) SetStatusCode(v int32) *DescribePolarxDataNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarxDataNodesResponse) SetBody(v *DescribePolarxDataNodesResponseBody) *DescribePolarxDataNodesResponse {
	s.Body = v
	return s
}

func (s *DescribePolarxDataNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
