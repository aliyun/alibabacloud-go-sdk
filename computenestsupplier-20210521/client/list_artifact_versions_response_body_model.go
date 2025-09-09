// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifacts(v []*ListArtifactVersionsResponseBodyArtifacts) *ListArtifactVersionsResponseBody
	GetArtifacts() []*ListArtifactVersionsResponseBodyArtifacts
	SetMaxResults(v int32) *ListArtifactVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListArtifactVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListArtifactVersionsResponseBody
	GetTotalCount() *int32
}

type ListArtifactVersionsResponseBody struct {
	// The version information about the deployment package.
	Artifacts []*ListArtifactVersionsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBody) GetArtifacts() []*ListArtifactVersionsResponseBodyArtifacts {
	return s.Artifacts
}

func (s *ListArtifactVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactVersionsResponseBody) SetArtifacts(v []*ListArtifactVersionsResponseBodyArtifacts) *ListArtifactVersionsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetMaxResults(v int32) *ListArtifactVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetNextToken(v string) *ListArtifactVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetRequestId(v string) *ListArtifactVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetTotalCount(v int32) *ListArtifactVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListArtifactVersionsResponseBodyArtifacts struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2022-10-20T02:19:53Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The distribution result of the image.
	ImageDelivery map[string]*string `json:"ImageDelivery,omitempty" xml:"ImageDelivery,omitempty"`
	// The distribution progress of the deployment package.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result file of the security scan.
	//
	// example:
	//
	// https://compute-nest-security-audit-bucket-ap-southeast-1.oss-ap-southeast-1.aliyuncs.com/51416747xxxx/xxxx
	ResultFile *string `json:"ResultFile,omitempty" xml:"ResultFile,omitempty"`
	// The result of the security scan. Valid values:
	//
	// 	- Normal: No risks exist on the deployment package.
	//
	// 	- AtRisk: Risks exist on the deployment package.
	//
	// 	- Processing: The deployment package is being scanned.
	//
	// example:
	//
	// Normal
	SecurityAuditResult *string `json:"SecurityAuditResult,omitempty" xml:"SecurityAuditResult,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- Created: The deployment package is created.
	//
	// 	- Scanning: The deployment package is being scanned.
	//
	// 	- ScanFailed: The deployment package failed to be scanned.
	//
	// 	- Delivering: The deployment package is being distributed.
	//
	// 	- Available: The deployment package is available.
	//
	// 	- Deleted: The deployment package is deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// [
	//
	// 					"cn-beijing",
	//
	// 					"cn-hangzhou",
	//
	// 					"cn-shanghai"
	//
	// 				]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListArtifactVersionsResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactBuildProperty() *string {
	return s.ArtifactBuildProperty
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactProperty() *string {
	return s.ArtifactProperty
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetImageDelivery() map[string]*string {
	return s.ImageDelivery
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetProgress() *string {
	return s.Progress
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetResultFile() *string {
	return s.ResultFile
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetSecurityAuditResult() *string {
	return s.SecurityAuditResult
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetStatus() *string {
	return s.Status
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetSupportRegionIds() *string {
	return s.SupportRegionIds
}

func (s *ListArtifactVersionsResponseBodyArtifacts) GetVersionName() *string {
	return s.VersionName
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactBuildProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactBuildType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactBuildType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactVersion(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtCreate(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtCreate = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetImageDelivery(v map[string]*string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ImageDelivery = v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetProgress(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Progress = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetResultFile(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ResultFile = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSecurityAuditResult(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SecurityAuditResult = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatus(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatusDetail(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.StatusDetail = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSupportRegionIds(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SupportRegionIds = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetVersionName(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.VersionName = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) Validate() error {
	return dara.Validate(s)
}
