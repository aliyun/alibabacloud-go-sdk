// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutOverReplicationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CutOverReplicationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CutOverReplicationJobResponse
	GetStatusCode() *int32
	SetBody(v *CutOverReplicationJobResponseBody) *CutOverReplicationJobResponse
	GetBody() *CutOverReplicationJobResponseBody
}

type CutOverReplicationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CutOverReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CutOverReplicationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CutOverReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CutOverReplicationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CutOverReplicationJobResponse) GetBody() *CutOverReplicationJobResponseBody {
	return s.Body
}

func (s *CutOverReplicationJobResponse) SetHeaders(v map[string]*string) *CutOverReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *CutOverReplicationJobResponse) SetStatusCode(v int32) *CutOverReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CutOverReplicationJobResponse) SetBody(v *CutOverReplicationJobResponseBody) *CutOverReplicationJobResponse {
	s.Body = v
	return s
}

func (s *CutOverReplicationJobResponse) Validate() error {
	return dara.Validate(s)
}
