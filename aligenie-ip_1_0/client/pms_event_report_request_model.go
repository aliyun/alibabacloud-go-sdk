// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPmsEventReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v string) *PmsEventReportRequest
	GetPayload() *string
}

type PmsEventReportRequest struct {
	// This parameter is required.
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s PmsEventReportRequest) String() string {
	return dara.Prettify(s)
}

func (s PmsEventReportRequest) GoString() string {
	return s.String()
}

func (s *PmsEventReportRequest) GetPayload() *string {
	return s.Payload
}

func (s *PmsEventReportRequest) SetPayload(v string) *PmsEventReportRequest {
	s.Payload = &v
	return s
}

func (s *PmsEventReportRequest) Validate() error {
	return dara.Validate(s)
}
