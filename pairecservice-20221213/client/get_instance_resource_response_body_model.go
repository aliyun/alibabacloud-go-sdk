// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetInstanceResourceResponseBody
	GetCategory() *string
	SetConfig(v string) *GetInstanceResourceResponseBody
	GetConfig() *string
	SetGmtCreateTime(v string) *GetInstanceResourceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceResourceResponseBody
	GetGmtModifiedTime() *string
	SetGroup(v string) *GetInstanceResourceResponseBody
	GetGroup() *string
	SetRequestId(v string) *GetInstanceResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *GetInstanceResourceResponseBody
	GetResourceId() *string
	SetType(v string) *GetInstanceResourceResponseBody
	GetType() *string
	SetUri(v string) *GetInstanceResourceResponseBody
	GetUri() *string
}

type GetInstanceResourceResponseBody struct {
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// D75C43DC-3D3A-5CC8-9AAC-8C77306C433B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetInstanceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceResponseBody) GetCategory() *string {
	return s.Category
}

func (s *GetInstanceResourceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetInstanceResourceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResourceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResourceResponseBody) GetGroup() *string {
	return s.Group
}

func (s *GetInstanceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetInstanceResourceResponseBody) GetType() *string {
	return s.Type
}

func (s *GetInstanceResourceResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetInstanceResourceResponseBody) SetCategory(v string) *GetInstanceResourceResponseBody {
	s.Category = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetConfig(v string) *GetInstanceResourceResponseBody {
	s.Config = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGmtCreateTime(v string) *GetInstanceResourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResourceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGroup(v string) *GetInstanceResourceResponseBody {
	s.Group = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetRequestId(v string) *GetInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetResourceId(v string) *GetInstanceResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetType(v string) *GetInstanceResourceResponseBody {
	s.Type = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetUri(v string) *GetInstanceResourceResponseBody {
	s.Uri = &v
	return s
}

func (s *GetInstanceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
