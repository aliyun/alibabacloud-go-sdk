// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionLiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MerchandisePlacementDetectionLiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MerchandisePlacementDetectionLiteResponse
	GetStatusCode() *int32
	SetBody(v *MerchandisePlacementDetectionLiteResponseBody) *MerchandisePlacementDetectionLiteResponse
	GetBody() *MerchandisePlacementDetectionLiteResponseBody
}

type MerchandisePlacementDetectionLiteResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MerchandisePlacementDetectionLiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MerchandisePlacementDetectionLiteResponse) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteResponse) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MerchandisePlacementDetectionLiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MerchandisePlacementDetectionLiteResponse) GetBody() *MerchandisePlacementDetectionLiteResponseBody {
	return s.Body
}

func (s *MerchandisePlacementDetectionLiteResponse) SetHeaders(v map[string]*string) *MerchandisePlacementDetectionLiteResponse {
	s.Headers = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponse) SetStatusCode(v int32) *MerchandisePlacementDetectionLiteResponse {
	s.StatusCode = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponse) SetBody(v *MerchandisePlacementDetectionLiteResponseBody) *MerchandisePlacementDetectionLiteResponse {
	s.Body = v
	return s
}

func (s *MerchandisePlacementDetectionLiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
