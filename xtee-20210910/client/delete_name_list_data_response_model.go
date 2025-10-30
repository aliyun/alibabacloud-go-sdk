// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNameListDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNameListDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNameListDataResponseBody) *DeleteNameListDataResponse
	GetBody() *DeleteNameListDataResponseBody
}

type DeleteNameListDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNameListDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNameListDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteNameListDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNameListDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNameListDataResponse) GetBody() *DeleteNameListDataResponseBody {
	return s.Body
}

func (s *DeleteNameListDataResponse) SetHeaders(v map[string]*string) *DeleteNameListDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteNameListDataResponse) SetStatusCode(v int32) *DeleteNameListDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNameListDataResponse) SetBody(v *DeleteNameListDataResponseBody) *DeleteNameListDataResponse {
	s.Body = v
	return s
}

func (s *DeleteNameListDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
