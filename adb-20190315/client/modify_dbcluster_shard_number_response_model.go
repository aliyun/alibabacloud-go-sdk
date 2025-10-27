// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterShardNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterShardNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterShardNumberResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterShardNumberResponseBody) *ModifyDBClusterShardNumberResponse
	GetBody() *ModifyDBClusterShardNumberResponseBody
}

type ModifyDBClusterShardNumberResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterShardNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterShardNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterShardNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterShardNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterShardNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterShardNumberResponse) GetBody() *ModifyDBClusterShardNumberResponseBody {
	return s.Body
}

func (s *ModifyDBClusterShardNumberResponse) SetHeaders(v map[string]*string) *ModifyDBClusterShardNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterShardNumberResponse) SetStatusCode(v int32) *ModifyDBClusterShardNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterShardNumberResponse) SetBody(v *ModifyDBClusterShardNumberResponseBody) *ModifyDBClusterShardNumberResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterShardNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
