// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDentriesInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDentriesInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDentriesInfoResponseBody) *QueryDentriesInfoResponse
	GetBody() *QueryDentriesInfoResponseBody
}

type QueryDentriesInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDentriesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDentriesInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDentriesInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDentriesInfoResponse) GetBody() *QueryDentriesInfoResponseBody {
	return s.Body
}

func (s *QueryDentriesInfoResponse) SetHeaders(v map[string]*string) *QueryDentriesInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDentriesInfoResponse) SetStatusCode(v int32) *QueryDentriesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDentriesInfoResponse) SetBody(v *QueryDentriesInfoResponseBody) *QueryDentriesInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDentriesInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
