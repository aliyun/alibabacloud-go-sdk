// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchAiboxesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureSearchAiboxesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureSearchAiboxesResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureSearchAiboxesResponseBody) *QueryPictureSearchAiboxesResponse
	GetBody() *QueryPictureSearchAiboxesResponseBody
}

type QueryPictureSearchAiboxesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchAiboxesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchAiboxesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureSearchAiboxesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureSearchAiboxesResponse) GetBody() *QueryPictureSearchAiboxesResponseBody {
	return s.Body
}

func (s *QueryPictureSearchAiboxesResponse) SetHeaders(v map[string]*string) *QueryPictureSearchAiboxesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchAiboxesResponse) SetStatusCode(v int32) *QueryPictureSearchAiboxesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponse) SetBody(v *QueryPictureSearchAiboxesResponseBody) *QueryPictureSearchAiboxesResponse {
	s.Body = v
	return s
}

func (s *QueryPictureSearchAiboxesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
