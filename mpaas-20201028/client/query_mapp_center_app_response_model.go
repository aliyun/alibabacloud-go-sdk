// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMappCenterAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMappCenterAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMappCenterAppResponse
	GetStatusCode() *int32
	SetBody(v *QueryMappCenterAppResponseBody) *QueryMappCenterAppResponse
	GetBody() *QueryMappCenterAppResponseBody
}

type QueryMappCenterAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMappCenterAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMappCenterAppResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponse) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMappCenterAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMappCenterAppResponse) GetBody() *QueryMappCenterAppResponseBody {
	return s.Body
}

func (s *QueryMappCenterAppResponse) SetHeaders(v map[string]*string) *QueryMappCenterAppResponse {
	s.Headers = v
	return s
}

func (s *QueryMappCenterAppResponse) SetStatusCode(v int32) *QueryMappCenterAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMappCenterAppResponse) SetBody(v *QueryMappCenterAppResponseBody) *QueryMappCenterAppResponse {
	s.Body = v
	return s
}

func (s *QueryMappCenterAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
