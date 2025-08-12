// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonSchemaRef(v []*GetUmodelResponseBodyCommonSchemaRef) *GetUmodelResponseBody
	GetCommonSchemaRef() []*GetUmodelResponseBodyCommonSchemaRef
	SetDescription(v string) *GetUmodelResponseBody
	GetDescription() *string
	SetRegionId(v string) *GetUmodelResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetUmodelResponseBody
	GetRequestId() *string
	SetWorkspace(v string) *GetUmodelResponseBody
	GetWorkspace() *string
}

type GetUmodelResponseBody struct {
	CommonSchemaRef []*GetUmodelResponseBodyCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 123-123123-sdf-435-3123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *GetUmodelResponseBody) GetCommonSchemaRef() []*GetUmodelResponseBodyCommonSchemaRef {
	return s.CommonSchemaRef
}

func (s *GetUmodelResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetUmodelResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUmodelResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetUmodelResponseBody) SetCommonSchemaRef(v []*GetUmodelResponseBodyCommonSchemaRef) *GetUmodelResponseBody {
	s.CommonSchemaRef = v
	return s
}

func (s *GetUmodelResponseBody) SetDescription(v string) *GetUmodelResponseBody {
	s.Description = &v
	return s
}

func (s *GetUmodelResponseBody) SetRegionId(v string) *GetUmodelResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetUmodelResponseBody) SetRequestId(v string) *GetUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUmodelResponseBody) SetWorkspace(v string) *GetUmodelResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetUmodelResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUmodelResponseBodyCommonSchemaRef struct {
	// example:
	//
	// test-job-123123
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// 5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetUmodelResponseBodyCommonSchemaRef) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelResponseBodyCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *GetUmodelResponseBodyCommonSchemaRef) GetGroup() *string {
	return s.Group
}

func (s *GetUmodelResponseBodyCommonSchemaRef) GetVersion() *string {
	return s.Version
}

func (s *GetUmodelResponseBodyCommonSchemaRef) SetGroup(v string) *GetUmodelResponseBodyCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *GetUmodelResponseBodyCommonSchemaRef) SetVersion(v string) *GetUmodelResponseBodyCommonSchemaRef {
	s.Version = &v
	return s
}

func (s *GetUmodelResponseBodyCommonSchemaRef) Validate() error {
	return dara.Validate(s)
}
