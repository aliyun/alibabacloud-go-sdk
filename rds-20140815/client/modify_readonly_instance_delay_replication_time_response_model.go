// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadonlyInstanceDelayReplicationTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReadonlyInstanceDelayReplicationTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReadonlyInstanceDelayReplicationTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) *ModifyReadonlyInstanceDelayReplicationTimeResponse
	GetBody() *ModifyReadonlyInstanceDelayReplicationTimeResponseBody
}

type ModifyReadonlyInstanceDelayReplicationTimeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReadonlyInstanceDelayReplicationTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReadonlyInstanceDelayReplicationTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadonlyInstanceDelayReplicationTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) GetBody() *ModifyReadonlyInstanceDelayReplicationTimeResponseBody {
	return s.Body
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) SetHeaders(v map[string]*string) *ModifyReadonlyInstanceDelayReplicationTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) SetStatusCode(v int32) *ModifyReadonlyInstanceDelayReplicationTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) SetBody(v *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) *ModifyReadonlyInstanceDelayReplicationTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
