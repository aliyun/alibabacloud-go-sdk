// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppTopologyResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppTopologyResponseBody) *QueryAppTopologyResponse
	GetBody() *QueryAppTopologyResponseBody
}

type QueryAppTopologyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppTopologyResponse) GoString() string {
	return s.String()
}

func (s *QueryAppTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppTopologyResponse) GetBody() *QueryAppTopologyResponseBody {
	return s.Body
}

func (s *QueryAppTopologyResponse) SetHeaders(v map[string]*string) *QueryAppTopologyResponse {
	s.Headers = v
	return s
}

func (s *QueryAppTopologyResponse) SetStatusCode(v int32) *QueryAppTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppTopologyResponse) SetBody(v *QueryAppTopologyResponseBody) *QueryAppTopologyResponse {
	s.Body = v
	return s
}

func (s *QueryAppTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
