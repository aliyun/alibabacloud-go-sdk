// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandiseSceneDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MerchandiseSceneDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MerchandiseSceneDetectionResponse
	GetStatusCode() *int32
	SetBody(v *MerchandiseSceneDetectionResponseBody) *MerchandiseSceneDetectionResponse
	GetBody() *MerchandiseSceneDetectionResponseBody
}

type MerchandiseSceneDetectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MerchandiseSceneDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MerchandiseSceneDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionResponse) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MerchandiseSceneDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MerchandiseSceneDetectionResponse) GetBody() *MerchandiseSceneDetectionResponseBody {
	return s.Body
}

func (s *MerchandiseSceneDetectionResponse) SetHeaders(v map[string]*string) *MerchandiseSceneDetectionResponse {
	s.Headers = v
	return s
}

func (s *MerchandiseSceneDetectionResponse) SetStatusCode(v int32) *MerchandiseSceneDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MerchandiseSceneDetectionResponse) SetBody(v *MerchandiseSceneDetectionResponseBody) *MerchandiseSceneDetectionResponse {
	s.Body = v
	return s
}

func (s *MerchandiseSceneDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
