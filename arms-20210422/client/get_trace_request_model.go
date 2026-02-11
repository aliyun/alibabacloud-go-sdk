// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetTraceRequest
	GetRegionId() *string
	SetTraceID(v string) *GetTraceRequest
	GetTraceID() *string
}

type GetTraceRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceRequest) GetTraceID() *string {
	return s.TraceID
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

func (s *GetTraceRequest) Validate() error {
	return dara.Validate(s)
}
