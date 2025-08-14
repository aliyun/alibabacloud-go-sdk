// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalDocStructureJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDigitalDocStructureJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDigitalDocStructureJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDigitalDocStructureJobResponseBody) *SubmitDigitalDocStructureJobResponse
	GetBody() *SubmitDigitalDocStructureJobResponseBody
}

type SubmitDigitalDocStructureJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDigitalDocStructureJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDigitalDocStructureJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalDocStructureJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDigitalDocStructureJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDigitalDocStructureJobResponse) GetBody() *SubmitDigitalDocStructureJobResponseBody {
	return s.Body
}

func (s *SubmitDigitalDocStructureJobResponse) SetHeaders(v map[string]*string) *SubmitDigitalDocStructureJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDigitalDocStructureJobResponse) SetStatusCode(v int32) *SubmitDigitalDocStructureJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponse) SetBody(v *SubmitDigitalDocStructureJobResponseBody) *SubmitDigitalDocStructureJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDigitalDocStructureJobResponse) Validate() error {
	return dara.Validate(s)
}
