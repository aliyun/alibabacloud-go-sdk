// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageCopyrightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImageCopyrightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImageCopyrightResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImageCopyrightResponseBody) *SubmitImageCopyrightResponse
	GetBody() *SubmitImageCopyrightResponseBody
}

type SubmitImageCopyrightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImageCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImageCopyrightResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageCopyrightResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImageCopyrightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImageCopyrightResponse) GetBody() *SubmitImageCopyrightResponseBody {
	return s.Body
}

func (s *SubmitImageCopyrightResponse) SetHeaders(v map[string]*string) *SubmitImageCopyrightResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageCopyrightResponse) SetStatusCode(v int32) *SubmitImageCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageCopyrightResponse) SetBody(v *SubmitImageCopyrightResponseBody) *SubmitImageCopyrightResponse {
	s.Body = v
	return s
}

func (s *SubmitImageCopyrightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
