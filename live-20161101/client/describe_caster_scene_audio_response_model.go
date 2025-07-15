// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterSceneAudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterSceneAudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterSceneAudioResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterSceneAudioResponseBody) *DescribeCasterSceneAudioResponse
	GetBody() *DescribeCasterSceneAudioResponseBody
}

type DescribeCasterSceneAudioResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterSceneAudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterSceneAudioResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterSceneAudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterSceneAudioResponse) GetBody() *DescribeCasterSceneAudioResponseBody {
	return s.Body
}

func (s *DescribeCasterSceneAudioResponse) SetHeaders(v map[string]*string) *DescribeCasterSceneAudioResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetStatusCode(v int32) *DescribeCasterSceneAudioResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetBody(v *DescribeCasterSceneAudioResponseBody) *DescribeCasterSceneAudioResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterSceneAudioResponse) Validate() error {
	return dara.Validate(s)
}
