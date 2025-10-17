// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaterializedViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaterializedViewResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaterializedViewResponseBody) *ModifyMaterializedViewResponse
	GetBody() *ModifyMaterializedViewResponseBody
}

type ModifyMaterializedViewResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaterializedViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaterializedViewResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaterializedViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaterializedViewResponse) GetBody() *ModifyMaterializedViewResponseBody {
	return s.Body
}

func (s *ModifyMaterializedViewResponse) SetHeaders(v map[string]*string) *ModifyMaterializedViewResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaterializedViewResponse) SetStatusCode(v int32) *ModifyMaterializedViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaterializedViewResponse) SetBody(v *ModifyMaterializedViewResponseBody) *ModifyMaterializedViewResponse {
	s.Body = v
	return s
}

func (s *ModifyMaterializedViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
