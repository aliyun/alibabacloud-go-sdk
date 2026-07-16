// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWaitingRoomPreviewPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWaitingRoomPreviewPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWaitingRoomPreviewPageResponse
	GetStatusCode() *int32
	SetBody(v *GetWaitingRoomPreviewPageResponseBody) *GetWaitingRoomPreviewPageResponse
	GetBody() *GetWaitingRoomPreviewPageResponseBody
}

type GetWaitingRoomPreviewPageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWaitingRoomPreviewPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWaitingRoomPreviewPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWaitingRoomPreviewPageResponse) GoString() string {
	return s.String()
}

func (s *GetWaitingRoomPreviewPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWaitingRoomPreviewPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWaitingRoomPreviewPageResponse) GetBody() *GetWaitingRoomPreviewPageResponseBody {
	return s.Body
}

func (s *GetWaitingRoomPreviewPageResponse) SetHeaders(v map[string]*string) *GetWaitingRoomPreviewPageResponse {
	s.Headers = v
	return s
}

func (s *GetWaitingRoomPreviewPageResponse) SetStatusCode(v int32) *GetWaitingRoomPreviewPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWaitingRoomPreviewPageResponse) SetBody(v *GetWaitingRoomPreviewPageResponseBody) *GetWaitingRoomPreviewPageResponse {
	s.Body = v
	return s
}

func (s *GetWaitingRoomPreviewPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
