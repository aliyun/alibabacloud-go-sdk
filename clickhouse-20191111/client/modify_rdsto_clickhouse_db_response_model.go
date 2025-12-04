// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRDSToClickhouseDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRDSToClickhouseDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRDSToClickhouseDbResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRDSToClickhouseDbResponseBody) *ModifyRDSToClickhouseDbResponse
	GetBody() *ModifyRDSToClickhouseDbResponseBody
}

type ModifyRDSToClickhouseDbResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRDSToClickhouseDbResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRDSToClickhouseDbResponse) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRDSToClickhouseDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRDSToClickhouseDbResponse) GetBody() *ModifyRDSToClickhouseDbResponseBody {
	return s.Body
}

func (s *ModifyRDSToClickhouseDbResponse) SetHeaders(v map[string]*string) *ModifyRDSToClickhouseDbResponse {
	s.Headers = v
	return s
}

func (s *ModifyRDSToClickhouseDbResponse) SetStatusCode(v int32) *ModifyRDSToClickhouseDbResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponse) SetBody(v *ModifyRDSToClickhouseDbResponseBody) *ModifyRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

func (s *ModifyRDSToClickhouseDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
