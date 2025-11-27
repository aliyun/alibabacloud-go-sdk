// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationDateClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLocationDateClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLocationDateClustersResponse
	GetStatusCode() *int32
	SetBody(v *QueryLocationDateClustersResponseBody) *QueryLocationDateClustersResponse
	GetBody() *QueryLocationDateClustersResponseBody
}

type QueryLocationDateClustersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLocationDateClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLocationDateClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationDateClustersResponse) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLocationDateClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLocationDateClustersResponse) GetBody() *QueryLocationDateClustersResponseBody {
	return s.Body
}

func (s *QueryLocationDateClustersResponse) SetHeaders(v map[string]*string) *QueryLocationDateClustersResponse {
	s.Headers = v
	return s
}

func (s *QueryLocationDateClustersResponse) SetStatusCode(v int32) *QueryLocationDateClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocationDateClustersResponse) SetBody(v *QueryLocationDateClustersResponseBody) *QueryLocationDateClustersResponse {
	s.Body = v
	return s
}

func (s *QueryLocationDateClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
