// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnnotationMissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAnnotationMissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAnnotationMissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAnnotationMissionResponseBody) *ModifyAnnotationMissionResponse
	GetBody() *ModifyAnnotationMissionResponseBody
}

type ModifyAnnotationMissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnnotationMissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAnnotationMissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnnotationMissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnnotationMissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAnnotationMissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAnnotationMissionResponse) GetBody() *ModifyAnnotationMissionResponseBody {
	return s.Body
}

func (s *ModifyAnnotationMissionResponse) SetHeaders(v map[string]*string) *ModifyAnnotationMissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnnotationMissionResponse) SetStatusCode(v int32) *ModifyAnnotationMissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnnotationMissionResponse) SetBody(v *ModifyAnnotationMissionResponseBody) *ModifyAnnotationMissionResponse {
	s.Body = v
	return s
}

func (s *ModifyAnnotationMissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
