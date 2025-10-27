// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIdList(v string) *GetVulStatisticsRequest
	GetGroupIdList() *string
	SetResourceDirectoryAccountId(v int64) *GetVulStatisticsRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *GetVulStatisticsRequest
	GetSourceIp() *string
	SetTypeList(v string) *GetVulStatisticsRequest
	GetTypeList() *string
}

type GetVulStatisticsRequest struct {
	// The ID of the asset group. Separate multiple IDs with commas (,).
	//
	// >  You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9997897
	GroupIdList *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 10.12.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the vulnerability whose statistics you want to query. Separate multiple types with commas (,). Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// 	- **app**: vulnerability detected by using a web scanner
	//
	// 	- **sca**: vulnerability detected based on software component analysis
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	TypeList *string `json:"TypeList,omitempty" xml:"TypeList,omitempty"`
}

func (s GetVulStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetVulStatisticsRequest) GetGroupIdList() *string {
	return s.GroupIdList
}

func (s *GetVulStatisticsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetVulStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetVulStatisticsRequest) GetTypeList() *string {
	return s.TypeList
}

func (s *GetVulStatisticsRequest) SetGroupIdList(v string) *GetVulStatisticsRequest {
	s.GroupIdList = &v
	return s
}

func (s *GetVulStatisticsRequest) SetResourceDirectoryAccountId(v int64) *GetVulStatisticsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetVulStatisticsRequest) SetSourceIp(v string) *GetVulStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetVulStatisticsRequest) SetTypeList(v string) *GetVulStatisticsRequest {
	s.TypeList = &v
	return s
}

func (s *GetVulStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
