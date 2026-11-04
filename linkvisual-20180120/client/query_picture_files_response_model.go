// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureFilesResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureFilesResponseBody) *QueryPictureFilesResponse
	GetBody() *QueryPictureFilesResponseBody
}

type QueryPictureFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureFilesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureFilesResponse) GetBody() *QueryPictureFilesResponseBody {
	return s.Body
}

func (s *QueryPictureFilesResponse) SetHeaders(v map[string]*string) *QueryPictureFilesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureFilesResponse) SetStatusCode(v int32) *QueryPictureFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureFilesResponse) SetBody(v *QueryPictureFilesResponseBody) *QueryPictureFilesResponse {
	s.Body = v
	return s
}

func (s *QueryPictureFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
