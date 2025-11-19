// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditSecurityIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAuditSecurityIpResponseBody
	GetRequestId() *string
	SetSecurityIpList(v []*ListAuditSecurityIpResponseBodySecurityIpList) *ListAuditSecurityIpResponseBody
	GetSecurityIpList() []*ListAuditSecurityIpResponseBodySecurityIpList
}

type ListAuditSecurityIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 664BBD08-C7DB-4E*****73-9D0958D9A899
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the review security group.
	SecurityIpList []*ListAuditSecurityIpResponseBodySecurityIpList `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty" type:"Repeated"`
}

func (s ListAuditSecurityIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuditSecurityIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuditSecurityIpResponseBody) GetSecurityIpList() []*ListAuditSecurityIpResponseBodySecurityIpList {
	return s.SecurityIpList
}

func (s *ListAuditSecurityIpResponseBody) SetRequestId(v string) *ListAuditSecurityIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuditSecurityIpResponseBody) SetSecurityIpList(v []*ListAuditSecurityIpResponseBodySecurityIpList) *ListAuditSecurityIpResponseBody {
	s.SecurityIpList = v
	return s
}

func (s *ListAuditSecurityIpResponseBody) Validate() error {
	if s.SecurityIpList != nil {
		for _, item := range s.SecurityIpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuditSecurityIpResponseBodySecurityIpList struct {
	// The time when the review security group was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-05-22T06:54:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IP addresses in the review security group.
	//
	// example:
	//
	// 30.27.14.0/24,30.39.127.245
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The time when the review security group was last modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-05-22T06:55:14Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The name of the review security group.
	//
	// example:
	//
	// Default
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ListAuditSecurityIpResponseBodySecurityIpList) String() string {
	return dara.Prettify(s)
}

func (s ListAuditSecurityIpResponseBodySecurityIpList) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) GetIps() *string {
	return s.Ips
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetCreationTime(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.CreationTime = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetIps(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.Ips = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetModificationTime(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.ModificationTime = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetSecurityGroupName(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.SecurityGroupName = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) Validate() error {
	return dara.Validate(s)
}
