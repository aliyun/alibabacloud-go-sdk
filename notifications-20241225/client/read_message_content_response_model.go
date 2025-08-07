// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMessageContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMessageContentResponse
	GetStatusCode() *int32
	SetBody(v *ReadMessageContentResponseBody) *ReadMessageContentResponse
	GetBody() *ReadMessageContentResponseBody
}

type ReadMessageContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMessageContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMessageContentResponse) GetBody() *ReadMessageContentResponseBody {
	return s.Body
}

func (s *ReadMessageContentResponse) SetHeaders(v map[string]*string) *ReadMessageContentResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageContentResponse) SetStatusCode(v int32) *ReadMessageContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageContentResponse) SetBody(v *ReadMessageContentResponseBody) *ReadMessageContentResponse {
	s.Body = v
	return s
}

func (s *ReadMessageContentResponse) Validate() error {
	return dara.Validate(s)
}
