// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocStructureJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocStructureJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocStructureJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocStructureJobResponseBody) *SubmitDocStructureJobResponse
	GetBody() *SubmitDocStructureJobResponseBody
}

type SubmitDocStructureJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocStructureJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocStructureJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocStructureJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocStructureJobResponse) GetBody() *SubmitDocStructureJobResponseBody {
	return s.Body
}

func (s *SubmitDocStructureJobResponse) SetHeaders(v map[string]*string) *SubmitDocStructureJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocStructureJobResponse) SetStatusCode(v int32) *SubmitDocStructureJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocStructureJobResponse) SetBody(v *SubmitDocStructureJobResponseBody) *SubmitDocStructureJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDocStructureJobResponse) Validate() error {
	return dara.Validate(s)
}
