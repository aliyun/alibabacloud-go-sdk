// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCartonCodeRephotographyDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MilkPowderCartonCodeRephotographyDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MilkPowderCartonCodeRephotographyDetectionResponse
	GetStatusCode() *int32
	SetBody(v *MilkPowderCartonCodeRephotographyDetectionResponseBody) *MilkPowderCartonCodeRephotographyDetectionResponse
	GetBody() *MilkPowderCartonCodeRephotographyDetectionResponseBody
}

type MilkPowderCartonCodeRephotographyDetectionResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MilkPowderCartonCodeRephotographyDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MilkPowderCartonCodeRephotographyDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionResponse) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) GetBody() *MilkPowderCartonCodeRephotographyDetectionResponseBody {
	return s.Body
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) SetHeaders(v map[string]*string) *MilkPowderCartonCodeRephotographyDetectionResponse {
	s.Headers = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) SetStatusCode(v int32) *MilkPowderCartonCodeRephotographyDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) SetBody(v *MilkPowderCartonCodeRephotographyDetectionResponseBody) *MilkPowderCartonCodeRephotographyDetectionResponse {
	s.Body = v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
