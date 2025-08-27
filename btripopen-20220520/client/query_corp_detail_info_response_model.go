// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCorpDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCorpDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCorpDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryCorpDetailInfoResponseBody) *QueryCorpDetailInfoResponse
	GetBody() *QueryCorpDetailInfoResponseBody
}

type QueryCorpDetailInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCorpDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCorpDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCorpDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCorpDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCorpDetailInfoResponse) GetBody() *QueryCorpDetailInfoResponseBody {
	return s.Body
}

func (s *QueryCorpDetailInfoResponse) SetHeaders(v map[string]*string) *QueryCorpDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpDetailInfoResponse) SetStatusCode(v int32) *QueryCorpDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCorpDetailInfoResponse) SetBody(v *QueryCorpDetailInfoResponseBody) *QueryCorpDetailInfoResponse {
	s.Body = v
	return s
}

func (s *QueryCorpDetailInfoResponse) Validate() error {
	return dara.Validate(s)
}
