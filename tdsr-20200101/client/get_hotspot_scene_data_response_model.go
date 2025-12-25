// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotSceneDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotSceneDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotSceneDataResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotSceneDataResponseBody) *GetHotspotSceneDataResponse
	GetBody() *GetHotspotSceneDataResponseBody
}

type GetHotspotSceneDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotSceneDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotSceneDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotSceneDataResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotSceneDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotSceneDataResponse) GetBody() *GetHotspotSceneDataResponseBody {
	return s.Body
}

func (s *GetHotspotSceneDataResponse) SetHeaders(v map[string]*string) *GetHotspotSceneDataResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotSceneDataResponse) SetStatusCode(v int32) *GetHotspotSceneDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotSceneDataResponse) SetBody(v *GetHotspotSceneDataResponseBody) *GetHotspotSceneDataResponse {
	s.Body = v
	return s
}

func (s *GetHotspotSceneDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
