// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotwordLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHotwordLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHotwordLibraryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHotwordLibraryResponseBody) *DeleteHotwordLibraryResponse
	GetBody() *DeleteHotwordLibraryResponseBody
}

type DeleteHotwordLibraryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotwordLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotwordLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotwordLibraryResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotwordLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHotwordLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotwordLibraryResponse) GetBody() *DeleteHotwordLibraryResponseBody {
	return s.Body
}

func (s *DeleteHotwordLibraryResponse) SetHeaders(v map[string]*string) *DeleteHotwordLibraryResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotwordLibraryResponse) SetStatusCode(v int32) *DeleteHotwordLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotwordLibraryResponse) SetBody(v *DeleteHotwordLibraryResponseBody) *DeleteHotwordLibraryResponse {
	s.Body = v
	return s
}

func (s *DeleteHotwordLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
