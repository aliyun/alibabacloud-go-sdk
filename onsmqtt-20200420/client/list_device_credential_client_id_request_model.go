// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCredentialClientIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListDeviceCredentialClientIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *ListDeviceCredentialClientIdRequest
	GetInstanceId() *string
	SetNextToken(v string) *ListDeviceCredentialClientIdRequest
	GetNextToken() *string
	SetPageNo(v string) *ListDeviceCredentialClientIdRequest
	GetPageNo() *string
	SetPageSize(v string) *ListDeviceCredentialClientIdRequest
	GetPageSize() *string
}

type ListDeviceCredentialClientIdRequest struct {
	// Group ID of the MQTT version of the micro message queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_xxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the Cloud Message Queue MQTT version instance, which must match the actual instance ID used by the client. You can obtain this ID from the **Instance Details*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Token for starting the next page query.
	//
	// example:
	//
	// FFdXXXXXWa
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the page number of the returned results. The starting page is counted from 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of query records to display per page.
	//
	// Value range: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceCredentialClientIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCredentialClientIdRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDeviceCredentialClientIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDeviceCredentialClientIdRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceCredentialClientIdRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListDeviceCredentialClientIdRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDeviceCredentialClientIdRequest) SetGroupId(v string) *ListDeviceCredentialClientIdRequest {
	s.GroupId = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetInstanceId(v string) *ListDeviceCredentialClientIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetNextToken(v string) *ListDeviceCredentialClientIdRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetPageNo(v string) *ListDeviceCredentialClientIdRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetPageSize(v string) *ListDeviceCredentialClientIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) Validate() error {
	return dara.Validate(s)
}
