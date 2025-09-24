// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetDatasetJobConfigResponseBody
	GetConfig() *string
	SetConfigType(v string) *GetDatasetJobConfigResponseBody
	GetConfigType() *string
	SetCreateTime(v string) *GetDatasetJobConfigResponseBody
	GetCreateTime() *string
	SetDatasetId(v string) *GetDatasetJobConfigResponseBody
	GetDatasetId() *string
	SetModifyTime(v string) *GetDatasetJobConfigResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *GetDatasetJobConfigResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetDatasetJobConfigResponseBody
	GetWorkspaceId() *string
}

type GetDatasetJobConfigResponseBody struct {
	// The configuration content. Configuration format for MultimodalIntelligentTag:
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	//
	// MultimodalSemanticIndex
	//
	// { "defaultModelId": "xxx" "defaultModelVersion":"1.0.0" }
	//
	// example:
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration type. Valid values:
	//
	// 	- MultimodalIntelligentTag
	//
	// 	- MultimodalSemanticIndex
	//
	// example:
	//
	// MultimodalIntelligentTag
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The time when the configuration is created.
	//
	// example:
	//
	// 2024-10-16T01:44:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d-lfd60v0p****ujtsdx
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The time when the configuration is modified.
	//
	// example:
	//
	// 2024-12-26T02:17:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D619B5C4B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 114243
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetJobConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetDatasetJobConfigResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetDatasetJobConfigResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDatasetJobConfigResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetDatasetJobConfigResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetDatasetJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetJobConfigResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetJobConfigResponseBody) SetConfig(v string) *GetDatasetJobConfigResponseBody {
	s.Config = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetConfigType(v string) *GetDatasetJobConfigResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetCreateTime(v string) *GetDatasetJobConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetDatasetId(v string) *GetDatasetJobConfigResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetModifyTime(v string) *GetDatasetJobConfigResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetRequestId(v string) *GetDatasetJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) SetWorkspaceId(v string) *GetDatasetJobConfigResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
