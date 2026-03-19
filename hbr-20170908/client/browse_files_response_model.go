// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowseFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BrowseFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BrowseFilesResponse
	GetStatusCode() *int32
	SetBody(v *BrowseFilesResponseBody) *BrowseFilesResponse
	GetBody() *BrowseFilesResponseBody
}

type BrowseFilesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BrowseFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BrowseFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s BrowseFilesResponse) GoString() string {
	return s.String()
}

func (s *BrowseFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BrowseFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BrowseFilesResponse) GetBody() *BrowseFilesResponseBody {
	return s.Body
}

func (s *BrowseFilesResponse) SetHeaders(v map[string]*string) *BrowseFilesResponse {
	s.Headers = v
	return s
}

func (s *BrowseFilesResponse) SetStatusCode(v int32) *BrowseFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *BrowseFilesResponse) SetBody(v *BrowseFilesResponseBody) *BrowseFilesResponse {
	s.Body = v
	return s
}

func (s *BrowseFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
