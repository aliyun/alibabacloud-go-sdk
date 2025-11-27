// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHADiagnoseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHADiagnoseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHADiagnoseConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHADiagnoseConfigResponseBody) *ModifyHADiagnoseConfigResponse
	GetBody() *ModifyHADiagnoseConfigResponseBody
}

type ModifyHADiagnoseConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHADiagnoseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHADiagnoseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHADiagnoseConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyHADiagnoseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHADiagnoseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHADiagnoseConfigResponse) GetBody() *ModifyHADiagnoseConfigResponseBody {
	return s.Body
}

func (s *ModifyHADiagnoseConfigResponse) SetHeaders(v map[string]*string) *ModifyHADiagnoseConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyHADiagnoseConfigResponse) SetStatusCode(v int32) *ModifyHADiagnoseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHADiagnoseConfigResponse) SetBody(v *ModifyHADiagnoseConfigResponseBody) *ModifyHADiagnoseConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyHADiagnoseConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
