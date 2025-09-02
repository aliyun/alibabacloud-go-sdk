// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicInfluenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicInfluenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicInfluenceResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicInfluenceResponseBody) *GetTopicInfluenceResponse
	GetBody() *GetTopicInfluenceResponseBody
}

type GetTopicInfluenceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicInfluenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicInfluenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicInfluenceResponse) GoString() string {
	return s.String()
}

func (s *GetTopicInfluenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicInfluenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicInfluenceResponse) GetBody() *GetTopicInfluenceResponseBody {
	return s.Body
}

func (s *GetTopicInfluenceResponse) SetHeaders(v map[string]*string) *GetTopicInfluenceResponse {
	s.Headers = v
	return s
}

func (s *GetTopicInfluenceResponse) SetStatusCode(v int32) *GetTopicInfluenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicInfluenceResponse) SetBody(v *GetTopicInfluenceResponseBody) *GetTopicInfluenceResponse {
	s.Body = v
	return s
}

func (s *GetTopicInfluenceResponse) Validate() error {
	return dara.Validate(s)
}
