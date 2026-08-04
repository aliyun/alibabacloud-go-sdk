// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOneKeyDeleteBlockListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOneKeyDeleteBlockListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOneKeyDeleteBlockListResponse
	GetStatusCode() *int32
	SetBody(v *QueryOneKeyDeleteBlockListResponseBody) *QueryOneKeyDeleteBlockListResponse
	GetBody() *QueryOneKeyDeleteBlockListResponseBody
}

type QueryOneKeyDeleteBlockListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOneKeyDeleteBlockListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOneKeyDeleteBlockListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOneKeyDeleteBlockListResponse) GoString() string {
	return s.String()
}

func (s *QueryOneKeyDeleteBlockListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOneKeyDeleteBlockListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOneKeyDeleteBlockListResponse) GetBody() *QueryOneKeyDeleteBlockListResponseBody {
	return s.Body
}

func (s *QueryOneKeyDeleteBlockListResponse) SetHeaders(v map[string]*string) *QueryOneKeyDeleteBlockListResponse {
	s.Headers = v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponse) SetStatusCode(v int32) *QueryOneKeyDeleteBlockListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponse) SetBody(v *QueryOneKeyDeleteBlockListResponseBody) *QueryOneKeyDeleteBlockListResponse {
	s.Body = v
	return s
}

func (s *QueryOneKeyDeleteBlockListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
