// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MerchandisePlacementDetectionProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MerchandisePlacementDetectionProResponse
	GetStatusCode() *int32
	SetBody(v *MerchandisePlacementDetectionProResponseBody) *MerchandisePlacementDetectionProResponse
	GetBody() *MerchandisePlacementDetectionProResponseBody
}

type MerchandisePlacementDetectionProResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MerchandisePlacementDetectionProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MerchandisePlacementDetectionProResponse) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionProResponse) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MerchandisePlacementDetectionProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MerchandisePlacementDetectionProResponse) GetBody() *MerchandisePlacementDetectionProResponseBody {
	return s.Body
}

func (s *MerchandisePlacementDetectionProResponse) SetHeaders(v map[string]*string) *MerchandisePlacementDetectionProResponse {
	s.Headers = v
	return s
}

func (s *MerchandisePlacementDetectionProResponse) SetStatusCode(v int32) *MerchandisePlacementDetectionProResponse {
	s.StatusCode = &v
	return s
}

func (s *MerchandisePlacementDetectionProResponse) SetBody(v *MerchandisePlacementDetectionProResponseBody) *MerchandisePlacementDetectionProResponse {
	s.Body = v
	return s
}

func (s *MerchandisePlacementDetectionProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
