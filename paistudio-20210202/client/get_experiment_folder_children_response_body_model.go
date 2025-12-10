// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentFolderChildrenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*GetExperimentFolderChildrenResponseBodyItems) *GetExperimentFolderChildrenResponseBody
	GetItems() []*GetExperimentFolderChildrenResponseBodyItems
	SetRequestId(v string) *GetExperimentFolderChildrenResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetExperimentFolderChildrenResponseBody
	GetTotalCount() *int32
}

type GetExperimentFolderChildrenResponseBody struct {
	Items []*GetExperimentFolderChildrenResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentFolderChildrenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentFolderChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponseBody) GetItems() []*GetExperimentFolderChildrenResponseBodyItems {
	return s.Items
}

func (s *GetExperimentFolderChildrenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentFolderChildrenResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetExperimentFolderChildrenResponseBody) SetItems(v []*GetExperimentFolderChildrenResponseBodyItems) *GetExperimentFolderChildrenResponseBody {
	s.Items = v
	return s
}

func (s *GetExperimentFolderChildrenResponseBody) SetRequestId(v string) *GetExperimentFolderChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBody) SetTotalCount(v int32) *GetExperimentFolderChildrenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExperimentFolderChildrenResponseBodyItems struct {
	// example:
	//
	// false
	Empty *bool `json:"Empty,omitempty" xml:"Empty,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// icon-folder
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// folder-xzf7t7****ch7qce
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dir
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentFolderChildrenResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentFolderChildrenResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetEmpty() *bool {
	return s.Empty
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetIcon() *string {
	return s.Icon
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *GetExperimentFolderChildrenResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetEmpty(v bool) *GetExperimentFolderChildrenResponseBodyItems {
	s.Empty = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetGmtCreateTime(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetGmtModifiedTime(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetIcon(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Icon = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetId(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Id = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetName(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Name = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetType(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Type = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
