// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPictureSearchJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPictureSearchJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPictureSearchJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPictureSearchJobStatusResponseBody) *GetPictureSearchJobStatusResponse
	GetBody() *GetPictureSearchJobStatusResponseBody
}

type GetPictureSearchJobStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPictureSearchJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPictureSearchJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPictureSearchJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPictureSearchJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPictureSearchJobStatusResponse) GetBody() *GetPictureSearchJobStatusResponseBody {
	return s.Body
}

func (s *GetPictureSearchJobStatusResponse) SetHeaders(v map[string]*string) *GetPictureSearchJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPictureSearchJobStatusResponse) SetStatusCode(v int32) *GetPictureSearchJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPictureSearchJobStatusResponse) SetBody(v *GetPictureSearchJobStatusResponseBody) *GetPictureSearchJobStatusResponse {
	s.Body = v
	return s
}

func (s *GetPictureSearchJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
