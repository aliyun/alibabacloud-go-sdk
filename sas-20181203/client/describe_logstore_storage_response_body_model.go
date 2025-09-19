// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogstoreStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DescribeLogstoreStorageResponseBody
	GetLogstore() *string
	SetPreserve(v int64) *DescribeLogstoreStorageResponseBody
	GetPreserve() *int64
	SetRequestId(v string) *DescribeLogstoreStorageResponseBody
	GetRequestId() *string
	SetTtl(v int32) *DescribeLogstoreStorageResponseBody
	GetTtl() *int32
	SetUsed(v int64) *DescribeLogstoreStorageResponseBody
	GetUsed() *int64
	SetUserProject(v string) *DescribeLogstoreStorageResponseBody
	GetUserProject() *string
}

type DescribeLogstoreStorageResponseBody struct {
	// The name of the dedicated Logstore that is used to store full logs of Security Center. The value is fixed as **sas-log**.
	//
	// example:
	//
	// sas-log
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The purchased log storage capacity, in GB.
	//
	// example:
	//
	// 12240
	Preserve *int64 `json:"Preserve,omitempty" xml:"Preserve,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 25EC270F-5783-4416-AD7C-1EDF063A039C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of days during which logs can be retained. The value is fixed as **180**, which indicates that logs can be retained for 180 days.
	//
	// >  You are not allowed to change the value of this parameter.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The used log storage capacity, in GB.
	//
	// example:
	//
	// 335
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
	// The name of the dedicated Project that is used to store full logs of Security Center.
	//
	// example:
	//
	// sas-log-XXXX-cn-hangzhou
	UserProject *string `json:"UserProject,omitempty" xml:"UserProject,omitempty"`
}

func (s DescribeLogstoreStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstoreStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogstoreStorageResponseBody) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeLogstoreStorageResponseBody) GetPreserve() *int64 {
	return s.Preserve
}

func (s *DescribeLogstoreStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogstoreStorageResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLogstoreStorageResponseBody) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeLogstoreStorageResponseBody) GetUserProject() *string {
	return s.UserProject
}

func (s *DescribeLogstoreStorageResponseBody) SetLogstore(v string) *DescribeLogstoreStorageResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) SetPreserve(v int64) *DescribeLogstoreStorageResponseBody {
	s.Preserve = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) SetRequestId(v string) *DescribeLogstoreStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) SetTtl(v int32) *DescribeLogstoreStorageResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) SetUsed(v int64) *DescribeLogstoreStorageResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) SetUserProject(v string) *DescribeLogstoreStorageResponseBody {
	s.UserProject = &v
	return s
}

func (s *DescribeLogstoreStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
