// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudGroupResponseBody) *ModifyHybridCloudGroupResponse
	GetBody() *ModifyHybridCloudGroupResponseBody
}

type ModifyHybridCloudGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudGroupResponse) GetBody() *ModifyHybridCloudGroupResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudGroupResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudGroupResponse) SetStatusCode(v int32) *ModifyHybridCloudGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudGroupResponse) SetBody(v *ModifyHybridCloudGroupResponseBody) *ModifyHybridCloudGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
