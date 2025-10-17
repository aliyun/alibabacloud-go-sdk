// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeHotReplicaModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeHotReplicaModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeHotReplicaModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeHotReplicaModeResponseBody) *ModifyDBNodeHotReplicaModeResponse
	GetBody() *ModifyDBNodeHotReplicaModeResponseBody
}

type ModifyDBNodeHotReplicaModeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeHotReplicaModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeHotReplicaModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeHotReplicaModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeHotReplicaModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeHotReplicaModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeHotReplicaModeResponse) GetBody() *ModifyDBNodeHotReplicaModeResponseBody {
	return s.Body
}

func (s *ModifyDBNodeHotReplicaModeResponse) SetHeaders(v map[string]*string) *ModifyDBNodeHotReplicaModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponse) SetStatusCode(v int32) *ModifyDBNodeHotReplicaModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponse) SetBody(v *ModifyDBNodeHotReplicaModeResponseBody) *ModifyDBNodeHotReplicaModeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
