// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetTraceAppRequest
	GetPid() *string
	SetRegionId(v string) *GetTraceAppRequest
	GetRegionId() *string
	SetTags(v []*GetTraceAppRequestTags) *GetTraceAppRequest
	GetTags() []*GetTraceAppRequestTags
}

type GetTraceAppRequest struct {
	// The process identifier (PID) of the application. For more information about how to obtain the PID, see [Obtain the PID of an application](https://www.alibabacloud.com/help/zh/doc-detail/186100.htm?spm=a2cdw.13409063.0.0.7a72281f0bkTfx#title-imy-7gj-qhr).
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*GetTraceAppRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetTraceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppRequest) GetPid() *string {
	return s.Pid
}

func (s *GetTraceAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceAppRequest) GetTags() []*GetTraceAppRequestTags {
	return s.Tags
}

func (s *GetTraceAppRequest) SetPid(v string) *GetTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *GetTraceAppRequest) SetRegionId(v string) *GetTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppRequest) SetTags(v []*GetTraceAppRequestTags) *GetTraceAppRequest {
	s.Tags = v
	return s
}

func (s *GetTraceAppRequest) Validate() error {
	return dara.Validate(s)
}

type GetTraceAppRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceAppRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppRequestTags) GoString() string {
	return s.String()
}

func (s *GetTraceAppRequestTags) GetKey() *string {
	return s.Key
}

func (s *GetTraceAppRequestTags) GetValue() *string {
	return s.Value
}

func (s *GetTraceAppRequestTags) SetKey(v string) *GetTraceAppRequestTags {
	s.Key = &v
	return s
}

func (s *GetTraceAppRequestTags) SetValue(v string) *GetTraceAppRequestTags {
	s.Value = &v
	return s
}

func (s *GetTraceAppRequestTags) Validate() error {
	return dara.Validate(s)
}
