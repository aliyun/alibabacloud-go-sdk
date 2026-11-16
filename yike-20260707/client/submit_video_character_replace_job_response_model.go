// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCharacterReplaceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoCharacterReplaceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoCharacterReplaceJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoCharacterReplaceJobResponseBody) *SubmitVideoCharacterReplaceJobResponse
	GetBody() *SubmitVideoCharacterReplaceJobResponseBody
}

type SubmitVideoCharacterReplaceJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoCharacterReplaceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoCharacterReplaceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCharacterReplaceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoCharacterReplaceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoCharacterReplaceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoCharacterReplaceJobResponse) GetBody() *SubmitVideoCharacterReplaceJobResponseBody {
	return s.Body
}

func (s *SubmitVideoCharacterReplaceJobResponse) SetHeaders(v map[string]*string) *SubmitVideoCharacterReplaceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoCharacterReplaceJobResponse) SetStatusCode(v int32) *SubmitVideoCharacterReplaceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobResponse) SetBody(v *SubmitVideoCharacterReplaceJobResponseBody) *SubmitVideoCharacterReplaceJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoCharacterReplaceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
