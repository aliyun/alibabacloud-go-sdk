// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureSearchAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureSearchAppsResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureSearchAppsResponseBody) *QueryPictureSearchAppsResponse
	GetBody() *QueryPictureSearchAppsResponseBody
}

type QueryPictureSearchAppsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAppsResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureSearchAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureSearchAppsResponse) GetBody() *QueryPictureSearchAppsResponseBody {
	return s.Body
}

func (s *QueryPictureSearchAppsResponse) SetHeaders(v map[string]*string) *QueryPictureSearchAppsResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchAppsResponse) SetStatusCode(v int32) *QueryPictureSearchAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchAppsResponse) SetBody(v *QueryPictureSearchAppsResponseBody) *QueryPictureSearchAppsResponse {
	s.Body = v
	return s
}

func (s *QueryPictureSearchAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
