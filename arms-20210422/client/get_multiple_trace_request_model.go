// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipleTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetMultipleTraceRequest
	GetRegionId() *string
	SetTraceIDs(v []*string) *GetMultipleTraceRequest
	GetTraceIDs() []*string
}

type GetMultipleTraceRequest struct {
	// This parameter is required.
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceIDs []*string `json:"TraceIDs,omitempty" xml:"TraceIDs,omitempty" type:"Repeated"`
}

func (s GetMultipleTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceRequest) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMultipleTraceRequest) GetTraceIDs() []*string {
	return s.TraceIDs
}

func (s *GetMultipleTraceRequest) SetRegionId(v string) *GetMultipleTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultipleTraceRequest) SetTraceIDs(v []*string) *GetMultipleTraceRequest {
	s.TraceIDs = v
	return s
}

func (s *GetMultipleTraceRequest) Validate() error {
	return dara.Validate(s)
}
