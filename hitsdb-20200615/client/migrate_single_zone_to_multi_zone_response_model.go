// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSingleZoneToMultiZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateSingleZoneToMultiZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateSingleZoneToMultiZoneResponse
	GetStatusCode() *int32
	SetBody(v *MigrateSingleZoneToMultiZoneResponseBody) *MigrateSingleZoneToMultiZoneResponse
	GetBody() *MigrateSingleZoneToMultiZoneResponseBody
}

type MigrateSingleZoneToMultiZoneResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateSingleZoneToMultiZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateSingleZoneToMultiZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateSingleZoneToMultiZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateSingleZoneToMultiZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateSingleZoneToMultiZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateSingleZoneToMultiZoneResponse) GetBody() *MigrateSingleZoneToMultiZoneResponseBody {
	return s.Body
}

func (s *MigrateSingleZoneToMultiZoneResponse) SetHeaders(v map[string]*string) *MigrateSingleZoneToMultiZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateSingleZoneToMultiZoneResponse) SetStatusCode(v int32) *MigrateSingleZoneToMultiZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneResponse) SetBody(v *MigrateSingleZoneToMultiZoneResponseBody) *MigrateSingleZoneToMultiZoneResponse {
	s.Body = v
	return s
}

func (s *MigrateSingleZoneToMultiZoneResponse) Validate() error {
	return dara.Validate(s)
}
