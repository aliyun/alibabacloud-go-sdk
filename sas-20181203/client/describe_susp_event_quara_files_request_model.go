// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventQuaraFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeSuspEventQuaraFilesRequest
	GetCurrentPage() *string
	SetFrom(v string) *DescribeSuspEventQuaraFilesRequest
	GetFrom() *string
	SetGroupId(v string) *DescribeSuspEventQuaraFilesRequest
	GetGroupId() *string
	SetGroupingId(v int64) *DescribeSuspEventQuaraFilesRequest
	GetGroupingId() *int64
	SetPageSize(v string) *DescribeSuspEventQuaraFilesRequest
	GetPageSize() *string
	SetQuaraTag(v string) *DescribeSuspEventQuaraFilesRequest
	GetQuaraTag() *string
	SetSourceIp(v string) *DescribeSuspEventQuaraFilesRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeSuspEventQuaraFilesRequest
	GetStatus() *string
}

type DescribeSuspEventQuaraFilesRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Deprecated
	//
	// The ID of the asset group.
	//
	// example:
	//
	// 10541428
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the server group to which the server belongs. The quarantined file is located on the server.
	//
	// example:
	//
	// 11472451
	GroupingId *int64 `json:"GroupingId,omitempty" xml:"GroupingId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID of the quarantined file.
	//
	// example:
	//
	// a31337789f64d39b2219733ec99f9af7
	QuaraTag *string `json:"QuaraTag,omitempty" xml:"QuaraTag,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 59.82.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the quarantined file that you want to query. Valid values:
	//
	// 	- **quaraFailed**: The file fails to be quarantined.
	//
	// 	- **quaraDone**: The file is quarantined.
	//
	// 	- **quaraing**: The file is being quarantined.
	//
	// 	- **rollbackFailed**: The system fails to cancel quarantining the file.
	//
	// 	- **rollbackDone**: The system cancelled quarantining the file.
	//
	// 	- **rollbacking**: The system is cancelling quarantining the file.
	//
	// example:
	//
	// quaraDone
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSuspEventQuaraFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSuspEventQuaraFilesRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventQuaraFilesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSuspEventQuaraFilesRequest) GetGroupingId() *int64 {
	return s.GroupingId
}

func (s *DescribeSuspEventQuaraFilesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSuspEventQuaraFilesRequest) GetQuaraTag() *string {
	return s.QuaraTag
}

func (s *DescribeSuspEventQuaraFilesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventQuaraFilesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSuspEventQuaraFilesRequest) SetCurrentPage(v string) *DescribeSuspEventQuaraFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetFrom(v string) *DescribeSuspEventQuaraFilesRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetGroupId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetGroupingId(v int64) *DescribeSuspEventQuaraFilesRequest {
	s.GroupingId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetPageSize(v string) *DescribeSuspEventQuaraFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetQuaraTag(v string) *DescribeSuspEventQuaraFilesRequest {
	s.QuaraTag = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetSourceIp(v string) *DescribeSuspEventQuaraFilesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetStatus(v string) *DescribeSuspEventQuaraFilesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) Validate() error {
	return dara.Validate(s)
}
