// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *GetDIJobRequest
	GetDIJobId() *int64
	SetWithDetails(v bool) *GetDIJobRequest
	GetWithDetails() *bool
}

type GetDIJobRequest struct {
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// Specifies whether to return detailed configuration information, including TransformationRules, TableMappings, and JobSettings. Valid values: true and false. Default value: true.
	//
	// example:
	//
	// true
	WithDetails *bool `json:"WithDetails,omitempty" xml:"WithDetails,omitempty"`
}

func (s GetDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *GetDIJobRequest) GetWithDetails() *bool {
	return s.WithDetails
}

func (s *GetDIJobRequest) SetDIJobId(v int64) *GetDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobRequest) SetWithDetails(v bool) *GetDIJobRequest {
	s.WithDetails = &v
	return s
}

func (s *GetDIJobRequest) Validate() error {
	return dara.Validate(s)
}
