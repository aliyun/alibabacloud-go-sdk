// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTranscodeTemplateGroupResponseBody) *DeleteTranscodeTemplateGroupResponse
	GetBody() *DeleteTranscodeTemplateGroupResponseBody
}

type DeleteTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTranscodeTemplateGroupResponse) GetBody() *DeleteTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *DeleteTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *DeleteTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponse) SetStatusCode(v int32) *DeleteTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponse) SetBody(v *DeleteTranscodeTemplateGroupResponseBody) *DeleteTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponse) Validate() error {
	return dara.Validate(s)
}
