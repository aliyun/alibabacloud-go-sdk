// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCanBottomCodeRephotographyDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MilkPowderCanBottomCodeRephotographyDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MilkPowderCanBottomCodeRephotographyDetectionResponse
	GetStatusCode() *int32
	SetBody(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) *MilkPowderCanBottomCodeRephotographyDetectionResponse
	GetBody() *MilkPowderCanBottomCodeRephotographyDetectionResponseBody
}

type MilkPowderCanBottomCodeRephotographyDetectionResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MilkPowderCanBottomCodeRephotographyDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionResponse) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) GetBody() *MilkPowderCanBottomCodeRephotographyDetectionResponseBody {
	return s.Body
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) SetHeaders(v map[string]*string) *MilkPowderCanBottomCodeRephotographyDetectionResponse {
	s.Headers = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) SetStatusCode(v int32) *MilkPowderCanBottomCodeRephotographyDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) SetBody(v *MilkPowderCanBottomCodeRephotographyDetectionResponseBody) *MilkPowderCanBottomCodeRephotographyDetectionResponse {
	s.Body = v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
