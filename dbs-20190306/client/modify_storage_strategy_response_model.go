// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStorageStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStorageStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStorageStrategyResponseBody) *ModifyStorageStrategyResponse
	GetBody() *ModifyStorageStrategyResponseBody
}

type ModifyStorageStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStorageStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStorageStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStorageStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStorageStrategyResponse) GetBody() *ModifyStorageStrategyResponseBody {
	return s.Body
}

func (s *ModifyStorageStrategyResponse) SetHeaders(v map[string]*string) *ModifyStorageStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyStorageStrategyResponse) SetStatusCode(v int32) *ModifyStorageStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStorageStrategyResponse) SetBody(v *ModifyStorageStrategyResponseBody) *ModifyStorageStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyStorageStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
