// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterAccessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterAccessWhiteListResponseBody) *ModifyDBClusterAccessWhiteListResponse
	GetBody() *ModifyDBClusterAccessWhiteListResponseBody
}

type ModifyDBClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterAccessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterAccessWhiteListResponse) GetBody() *ModifyDBClusterAccessWhiteListResponseBody {
	return s.Body
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetStatusCode(v int32) *ModifyDBClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetBody(v *ModifyDBClusterAccessWhiteListResponseBody) *ModifyDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
