// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCharacterReplaceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoCharacterReplaceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoCharacterReplaceJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoCharacterReplaceJobResponseBody) *GetVideoCharacterReplaceJobResponse
	GetBody() *GetVideoCharacterReplaceJobResponseBody
}

type GetVideoCharacterReplaceJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoCharacterReplaceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoCharacterReplaceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoCharacterReplaceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoCharacterReplaceJobResponse) GetBody() *GetVideoCharacterReplaceJobResponseBody {
	return s.Body
}

func (s *GetVideoCharacterReplaceJobResponse) SetHeaders(v map[string]*string) *GetVideoCharacterReplaceJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponse) SetStatusCode(v int32) *GetVideoCharacterReplaceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponse) SetBody(v *GetVideoCharacterReplaceJobResponseBody) *GetVideoCharacterReplaceJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
