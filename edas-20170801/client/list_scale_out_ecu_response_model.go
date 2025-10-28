// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScaleOutEcuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScaleOutEcuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScaleOutEcuResponse
	GetStatusCode() *int32
	SetBody(v *ListScaleOutEcuResponseBody) *ListScaleOutEcuResponse
	GetBody() *ListScaleOutEcuResponseBody
}

type ListScaleOutEcuResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScaleOutEcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScaleOutEcuResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScaleOutEcuResponse) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScaleOutEcuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScaleOutEcuResponse) GetBody() *ListScaleOutEcuResponseBody {
	return s.Body
}

func (s *ListScaleOutEcuResponse) SetHeaders(v map[string]*string) *ListScaleOutEcuResponse {
	s.Headers = v
	return s
}

func (s *ListScaleOutEcuResponse) SetStatusCode(v int32) *ListScaleOutEcuResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScaleOutEcuResponse) SetBody(v *ListScaleOutEcuResponseBody) *ListScaleOutEcuResponse {
	s.Body = v
	return s
}

func (s *ListScaleOutEcuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
