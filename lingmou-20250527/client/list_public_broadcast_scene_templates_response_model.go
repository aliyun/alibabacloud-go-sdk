// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicBroadcastSceneTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicBroadcastSceneTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicBroadcastSceneTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicBroadcastSceneTemplatesResponseBody) *ListPublicBroadcastSceneTemplatesResponse
	GetBody() *ListPublicBroadcastSceneTemplatesResponseBody
}

type ListPublicBroadcastSceneTemplatesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicBroadcastSceneTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicBroadcastSceneTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicBroadcastSceneTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPublicBroadcastSceneTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicBroadcastSceneTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicBroadcastSceneTemplatesResponse) GetBody() *ListPublicBroadcastSceneTemplatesResponseBody {
	return s.Body
}

func (s *ListPublicBroadcastSceneTemplatesResponse) SetHeaders(v map[string]*string) *ListPublicBroadcastSceneTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponse) SetStatusCode(v int32) *ListPublicBroadcastSceneTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponse) SetBody(v *ListPublicBroadcastSceneTemplatesResponseBody) *ListPublicBroadcastSceneTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
