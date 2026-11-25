// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManifestEditTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListManifestEditTemplateResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListManifestEditTemplateResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListManifestEditTemplateResponseBody
	GetRequestId() *string
	SetTemplateList(v []*ListManifestEditTemplateResponseBodyTemplateList) *ListManifestEditTemplateResponseBody
	GetTemplateList() []*ListManifestEditTemplateResponseBodyTemplateList
	SetTotalCount(v int32) *ListManifestEditTemplateResponseBody
	GetTotalCount() *int32
}

type ListManifestEditTemplateResponseBody struct {
	// The maximum number of entries returned for the current request.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position returned by the current call. An empty value indicates that all data has been retrieved.
	//
	// This parameter is required.
	//
	// example:
	//
	// page_2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of ManifestEdit templates.
	TemplateList []*ListManifestEditTemplateResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	// The total number of entries that meet the current request conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListManifestEditTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManifestEditTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListManifestEditTemplateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManifestEditTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManifestEditTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManifestEditTemplateResponseBody) GetTemplateList() []*ListManifestEditTemplateResponseBodyTemplateList {
	return s.TemplateList
}

func (s *ListManifestEditTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListManifestEditTemplateResponseBody) SetMaxResults(v int32) *ListManifestEditTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListManifestEditTemplateResponseBody) SetNextToken(v string) *ListManifestEditTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManifestEditTemplateResponseBody) SetRequestId(v string) *ListManifestEditTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManifestEditTemplateResponseBody) SetTemplateList(v []*ListManifestEditTemplateResponseBodyTemplateList) *ListManifestEditTemplateResponseBody {
	s.TemplateList = v
	return s
}

func (s *ListManifestEditTemplateResponseBody) SetTotalCount(v int32) *ListManifestEditTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListManifestEditTemplateResponseBody) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListManifestEditTemplateResponseBodyTemplateList struct {
	// The application ID. Default value: **app-1000000**. For more information, see [Multiple applications](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The creation time. The time follows the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2026-04-10T08:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The modification time. The time follows the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2026-04-10T08:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the track template group.
	//
	// example:
	//
	// multilang-tracks
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template status. A value of **1*	- indicates enabled, and a value of **0*	- indicates disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ManifestEdit template ID.
	//
	// example:
	//
	// 9c6581d59240de5e258ffa4c71a3e01b
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListManifestEditTemplateResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListManifestEditTemplateResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetAppId() *string {
	return s.AppId
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetName() *string {
	return s.Name
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetStatus() *int32 {
	return s.Status
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetAppId(v string) *ListManifestEditTemplateResponseBodyTemplateList {
	s.AppId = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetCreationTime(v string) *ListManifestEditTemplateResponseBodyTemplateList {
	s.CreationTime = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetModifyTime(v string) *ListManifestEditTemplateResponseBodyTemplateList {
	s.ModifyTime = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetName(v string) *ListManifestEditTemplateResponseBodyTemplateList {
	s.Name = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetStatus(v int32) *ListManifestEditTemplateResponseBodyTemplateList {
	s.Status = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) SetTemplateId(v string) *ListManifestEditTemplateResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListManifestEditTemplateResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
