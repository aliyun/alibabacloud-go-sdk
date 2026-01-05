// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *DescribeComponentResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *DescribeComponentResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *DescribeComponentResponseBody
	GetName() *string
	SetOwnerId(v string) *DescribeComponentResponseBody
	GetOwnerId() *string
	SetRenderedContent(v string) *DescribeComponentResponseBody
	GetRenderedContent() *string
	SetRequestId(v string) *DescribeComponentResponseBody
	GetRequestId() *string
	SetTemplate(v *DescribeComponentResponseBodyTemplate) *DescribeComponentResponseBody
	GetTemplate() *DescribeComponentResponseBodyTemplate
	SetUserId(v string) *DescribeComponentResponseBody
	GetUserId() *string
	SetUuid(v string) *DescribeComponentResponseBody
	GetUuid() *string
	SetVersion(v string) *DescribeComponentResponseBody
	GetVersion() *string
}

type DescribeComponentResponseBody struct {
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// dataset-accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ---\\n# Source: dataset-accelerator/templates/dataset-accelerator\\napiVersion: data.datasetacc.io/v1\\nkind: DataSetAccelerator\\nmetadata:\\n  name: dataset-accelerator\\nspec:\\n  frontEndAddr: 10.0.0.2\\n  frontEndPort: 7001\\n
	RenderedContent *string `json:"RenderedContent,omitempty" xml:"RenderedContent,omitempty"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *DescribeComponentResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// cmpt-5zk866779me51jgu3w
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// v1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DescribeComponentResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *DescribeComponentResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeComponentResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeComponentResponseBody) GetRenderedContent() *string {
	return s.RenderedContent
}

func (s *DescribeComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentResponseBody) GetTemplate() *DescribeComponentResponseBodyTemplate {
	return s.Template
}

func (s *DescribeComponentResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeComponentResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeComponentResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeComponentResponseBody) SetGmtCreateTime(v string) *DescribeComponentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeComponentResponseBody) SetGmtModifiedTime(v string) *DescribeComponentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeComponentResponseBody) SetName(v string) *DescribeComponentResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeComponentResponseBody) SetOwnerId(v string) *DescribeComponentResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetRenderedContent(v string) *DescribeComponentResponseBody {
	s.RenderedContent = &v
	return s
}

func (s *DescribeComponentResponseBody) SetRequestId(v string) *DescribeComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetTemplate(v *DescribeComponentResponseBodyTemplate) *DescribeComponentResponseBody {
	s.Template = v
	return s
}

func (s *DescribeComponentResponseBody) SetUserId(v string) *DescribeComponentResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetUuid(v string) *DescribeComponentResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeComponentResponseBody) SetVersion(v string) *DescribeComponentResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeComponentResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeComponentResponseBodyTemplate struct {
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// data/VOCdevkit/VOC2007/ImageSets/Main/val.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeComponentResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponseBodyTemplate) GetType() *string {
	return s.Type
}

func (s *DescribeComponentResponseBodyTemplate) GetUri() *string {
	return s.Uri
}

func (s *DescribeComponentResponseBodyTemplate) SetType(v string) *DescribeComponentResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *DescribeComponentResponseBodyTemplate) SetUri(v string) *DescribeComponentResponseBodyTemplate {
	s.Uri = &v
	return s
}

func (s *DescribeComponentResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
