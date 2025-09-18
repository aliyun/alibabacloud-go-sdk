// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartHomeSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartHomeSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartHomeSceneListResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartHomeSceneListResponseBody) *GetSmartHomeSceneListResponse
	GetBody() *GetSmartHomeSceneListResponseBody
}

type GetSmartHomeSceneListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartHomeSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartHomeSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHomeSceneListResponse) GoString() string {
	return s.String()
}

func (s *GetSmartHomeSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartHomeSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartHomeSceneListResponse) GetBody() *GetSmartHomeSceneListResponseBody {
	return s.Body
}

func (s *GetSmartHomeSceneListResponse) SetHeaders(v map[string]*string) *GetSmartHomeSceneListResponse {
	s.Headers = v
	return s
}

func (s *GetSmartHomeSceneListResponse) SetStatusCode(v int32) *GetSmartHomeSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartHomeSceneListResponse) SetBody(v *GetSmartHomeSceneListResponseBody) *GetSmartHomeSceneListResponse {
	s.Body = v
	return s
}

func (s *GetSmartHomeSceneListResponse) Validate() error {
	return dara.Validate(s)
}
