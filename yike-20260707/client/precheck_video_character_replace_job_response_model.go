// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckVideoCharacterReplaceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckVideoCharacterReplaceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckVideoCharacterReplaceJobResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckVideoCharacterReplaceJobResponseBody) *PrecheckVideoCharacterReplaceJobResponse
	GetBody() *PrecheckVideoCharacterReplaceJobResponseBody
}

type PrecheckVideoCharacterReplaceJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckVideoCharacterReplaceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckVideoCharacterReplaceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckVideoCharacterReplaceJobResponse) GoString() string {
	return s.String()
}

func (s *PrecheckVideoCharacterReplaceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckVideoCharacterReplaceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckVideoCharacterReplaceJobResponse) GetBody() *PrecheckVideoCharacterReplaceJobResponseBody {
	return s.Body
}

func (s *PrecheckVideoCharacterReplaceJobResponse) SetHeaders(v map[string]*string) *PrecheckVideoCharacterReplaceJobResponse {
	s.Headers = v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponse) SetStatusCode(v int32) *PrecheckVideoCharacterReplaceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponse) SetBody(v *PrecheckVideoCharacterReplaceJobResponseBody) *PrecheckVideoCharacterReplaceJobResponse {
	s.Body = v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
