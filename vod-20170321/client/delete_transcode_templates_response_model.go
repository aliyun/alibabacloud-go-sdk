// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTranscodeTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTranscodeTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTranscodeTemplatesResponseBody) *DeleteTranscodeTemplatesResponse
	GetBody() *DeleteTranscodeTemplatesResponseBody
}

type DeleteTranscodeTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTranscodeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTranscodeTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTranscodeTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTranscodeTemplatesResponse) GetBody() *DeleteTranscodeTemplatesResponseBody {
	return s.Body
}

func (s *DeleteTranscodeTemplatesResponse) SetHeaders(v map[string]*string) *DeleteTranscodeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTranscodeTemplatesResponse) SetStatusCode(v int32) *DeleteTranscodeTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTranscodeTemplatesResponse) SetBody(v *DeleteTranscodeTemplatesResponseBody) *DeleteTranscodeTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteTranscodeTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
