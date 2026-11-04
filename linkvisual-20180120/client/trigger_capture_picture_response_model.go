// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCapturePictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerCapturePictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerCapturePictureResponse
	GetStatusCode() *int32
	SetBody(v *TriggerCapturePictureResponseBody) *TriggerCapturePictureResponse
	GetBody() *TriggerCapturePictureResponseBody
}

type TriggerCapturePictureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerCapturePictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerCapturePictureResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerCapturePictureResponse) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerCapturePictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerCapturePictureResponse) GetBody() *TriggerCapturePictureResponseBody {
	return s.Body
}

func (s *TriggerCapturePictureResponse) SetHeaders(v map[string]*string) *TriggerCapturePictureResponse {
	s.Headers = v
	return s
}

func (s *TriggerCapturePictureResponse) SetStatusCode(v int32) *TriggerCapturePictureResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerCapturePictureResponse) SetBody(v *TriggerCapturePictureResponseBody) *TriggerCapturePictureResponse {
	s.Body = v
	return s
}

func (s *TriggerCapturePictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
