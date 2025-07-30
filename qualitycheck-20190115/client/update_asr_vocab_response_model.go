// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAsrVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAsrVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAsrVocabResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAsrVocabResponseBody) *UpdateAsrVocabResponse
	GetBody() *UpdateAsrVocabResponseBody
}

type UpdateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAsrVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAsrVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAsrVocabResponse) GetBody() *UpdateAsrVocabResponseBody {
	return s.Body
}

func (s *UpdateAsrVocabResponse) SetHeaders(v map[string]*string) *UpdateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *UpdateAsrVocabResponse) SetStatusCode(v int32) *UpdateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAsrVocabResponse) SetBody(v *UpdateAsrVocabResponseBody) *UpdateAsrVocabResponse {
	s.Body = v
	return s
}

func (s *UpdateAsrVocabResponse) Validate() error {
	return dara.Validate(s)
}
