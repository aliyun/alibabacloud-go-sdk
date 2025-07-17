// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDAGVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagVersionList(v *ListDAGVersionsResponseBodyDagVersionList) *ListDAGVersionsResponseBody
	GetDagVersionList() *ListDAGVersionsResponseBodyDagVersionList
	SetErrorCode(v string) *ListDAGVersionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDAGVersionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDAGVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDAGVersionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDAGVersionsResponseBody
	GetTotalCount() *int32
}

type ListDAGVersionsResponseBody struct {
	// The information about the published versions.
	DagVersionList *ListDAGVersionsResponseBodyDagVersionList `json:"DagVersionList,omitempty" xml:"DagVersionList,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7775630-7901-51B9-8782-9B585EC0799A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDAGVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDAGVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDAGVersionsResponseBody) GetDagVersionList() *ListDAGVersionsResponseBodyDagVersionList {
	return s.DagVersionList
}

func (s *ListDAGVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDAGVersionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDAGVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDAGVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDAGVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDAGVersionsResponseBody) SetDagVersionList(v *ListDAGVersionsResponseBodyDagVersionList) *ListDAGVersionsResponseBody {
	s.DagVersionList = v
	return s
}

func (s *ListDAGVersionsResponseBody) SetErrorCode(v string) *ListDAGVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDAGVersionsResponseBody) SetErrorMessage(v string) *ListDAGVersionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDAGVersionsResponseBody) SetRequestId(v string) *ListDAGVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDAGVersionsResponseBody) SetSuccess(v bool) *ListDAGVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDAGVersionsResponseBody) SetTotalCount(v int32) *ListDAGVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDAGVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDAGVersionsResponseBodyDagVersionList struct {
	DagVersion []*ListDAGVersionsResponseBodyDagVersionListDagVersion `json:"DagVersion,omitempty" xml:"DagVersion,omitempty" type:"Repeated"`
}

func (s ListDAGVersionsResponseBodyDagVersionList) String() string {
	return dara.Prettify(s)
}

func (s ListDAGVersionsResponseBodyDagVersionList) GoString() string {
	return s.String()
}

func (s *ListDAGVersionsResponseBodyDagVersionList) GetDagVersion() []*ListDAGVersionsResponseBodyDagVersionListDagVersion {
	return s.DagVersion
}

func (s *ListDAGVersionsResponseBodyDagVersionList) SetDagVersion(v []*ListDAGVersionsResponseBodyDagVersionListDagVersion) *ListDAGVersionsResponseBodyDagVersionList {
	s.DagVersion = v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionList) Validate() error {
	return dara.Validate(s)
}

type ListDAGVersionsResponseBodyDagVersionListDagVersion struct {
	// The name of the task flow.
	//
	// example:
	//
	// Spark SQL
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The ID of the task flow owner.
	//
	// example:
	//
	// 51****
	DagOwnerId *string `json:"DagOwnerId,omitempty" xml:"DagOwnerId,omitempty"`
	// The name of the task flow owner.
	//
	// example:
	//
	// name
	DagOwnerNickName *string `json:"DagOwnerNickName,omitempty" xml:"DagOwnerNickName,omitempty"`
	// The ID of the previously published version.
	//
	// example:
	//
	// 2****
	LastVersionId *int64 `json:"LastVersionId,omitempty" xml:"LastVersionId,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// test_OSS
	VersionComments *string `json:"VersionComments,omitempty" xml:"VersionComments,omitempty"`
	// The ID of the version.
	//
	// example:
	//
	// 2****
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListDAGVersionsResponseBodyDagVersionListDagVersion) String() string {
	return dara.Prettify(s)
}

func (s ListDAGVersionsResponseBodyDagVersionListDagVersion) GoString() string {
	return s.String()
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetDagName() *string {
	return s.DagName
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetDagOwnerId() *string {
	return s.DagOwnerId
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetDagOwnerNickName() *string {
	return s.DagOwnerNickName
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetLastVersionId() *int64 {
	return s.LastVersionId
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetVersionComments() *string {
	return s.VersionComments
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) GetVersionId() *int64 {
	return s.VersionId
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetDagName(v string) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.DagName = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetDagOwnerId(v string) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.DagOwnerId = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetDagOwnerNickName(v string) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.DagOwnerNickName = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetLastVersionId(v int64) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.LastVersionId = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetVersionComments(v string) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.VersionComments = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) SetVersionId(v int64) *ListDAGVersionsResponseBodyDagVersionListDagVersion {
	s.VersionId = &v
	return s
}

func (s *ListDAGVersionsResponseBodyDagVersionListDagVersion) Validate() error {
	return dara.Validate(s)
}
