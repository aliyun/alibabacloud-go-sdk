// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVocabResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVocabResponseBody) *DeleteVocabResponse
	GetBody() *DeleteVocabResponseBody
}

type DeleteVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabResponse) GoString() string {
	return s.String()
}

func (s *DeleteVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVocabResponse) GetBody() *DeleteVocabResponseBody {
	return s.Body
}

func (s *DeleteVocabResponse) SetHeaders(v map[string]*string) *DeleteVocabResponse {
	s.Headers = v
	return s
}

func (s *DeleteVocabResponse) SetStatusCode(v int32) *DeleteVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVocabResponse) SetBody(v *DeleteVocabResponseBody) *DeleteVocabResponse {
	s.Body = v
	return s
}

func (s *DeleteVocabResponse) Validate() error {
	return dara.Validate(s)
}
