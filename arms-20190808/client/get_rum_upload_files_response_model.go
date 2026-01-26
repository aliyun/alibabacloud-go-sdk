// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumUploadFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumUploadFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumUploadFilesResponse
	GetStatusCode() *int32
	SetBody(v *GetRumUploadFilesResponseBody) *GetRumUploadFilesResponse
	GetBody() *GetRumUploadFilesResponseBody
}

type GetRumUploadFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumUploadFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumUploadFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumUploadFilesResponse) GoString() string {
	return s.String()
}

func (s *GetRumUploadFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumUploadFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumUploadFilesResponse) GetBody() *GetRumUploadFilesResponseBody {
	return s.Body
}

func (s *GetRumUploadFilesResponse) SetHeaders(v map[string]*string) *GetRumUploadFilesResponse {
	s.Headers = v
	return s
}

func (s *GetRumUploadFilesResponse) SetStatusCode(v int32) *GetRumUploadFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumUploadFilesResponse) SetBody(v *GetRumUploadFilesResponseBody) *GetRumUploadFilesResponse {
	s.Body = v
	return s
}

func (s *GetRumUploadFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
