// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MerchandisePlacementDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MerchandisePlacementDetectionResponse
	GetStatusCode() *int32
	SetBody(v *MerchandisePlacementDetectionResponseBody) *MerchandisePlacementDetectionResponse
	GetBody() *MerchandisePlacementDetectionResponseBody
}

type MerchandisePlacementDetectionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MerchandisePlacementDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MerchandisePlacementDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionResponse) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MerchandisePlacementDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MerchandisePlacementDetectionResponse) GetBody() *MerchandisePlacementDetectionResponseBody {
	return s.Body
}

func (s *MerchandisePlacementDetectionResponse) SetHeaders(v map[string]*string) *MerchandisePlacementDetectionResponse {
	s.Headers = v
	return s
}

func (s *MerchandisePlacementDetectionResponse) SetStatusCode(v int32) *MerchandisePlacementDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MerchandisePlacementDetectionResponse) SetBody(v *MerchandisePlacementDetectionResponseBody) *MerchandisePlacementDetectionResponse {
	s.Body = v
	return s
}

func (s *MerchandisePlacementDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
