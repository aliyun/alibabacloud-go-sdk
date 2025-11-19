// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWuyingServerSceneUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateWuyingServerSceneUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateWuyingServerSceneUrlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateWuyingServerSceneUrlResponseBody) *GenerateWuyingServerSceneUrlResponse
	GetBody() *GenerateWuyingServerSceneUrlResponseBody
}

type GenerateWuyingServerSceneUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateWuyingServerSceneUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateWuyingServerSceneUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateWuyingServerSceneUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateWuyingServerSceneUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateWuyingServerSceneUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateWuyingServerSceneUrlResponse) GetBody() *GenerateWuyingServerSceneUrlResponseBody {
	return s.Body
}

func (s *GenerateWuyingServerSceneUrlResponse) SetHeaders(v map[string]*string) *GenerateWuyingServerSceneUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponse) SetStatusCode(v int32) *GenerateWuyingServerSceneUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponse) SetBody(v *GenerateWuyingServerSceneUrlResponseBody) *GenerateWuyingServerSceneUrlResponse {
	s.Body = v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
