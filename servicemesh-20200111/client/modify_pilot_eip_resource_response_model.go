// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPilotEipResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPilotEipResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPilotEipResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPilotEipResourceResponseBody) *ModifyPilotEipResourceResponse
	GetBody() *ModifyPilotEipResourceResponseBody
}

type ModifyPilotEipResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPilotEipResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPilotEipResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPilotEipResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyPilotEipResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPilotEipResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPilotEipResourceResponse) GetBody() *ModifyPilotEipResourceResponseBody {
	return s.Body
}

func (s *ModifyPilotEipResourceResponse) SetHeaders(v map[string]*string) *ModifyPilotEipResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyPilotEipResourceResponse) SetStatusCode(v int32) *ModifyPilotEipResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPilotEipResourceResponse) SetBody(v *ModifyPilotEipResourceResponseBody) *ModifyPilotEipResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyPilotEipResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
