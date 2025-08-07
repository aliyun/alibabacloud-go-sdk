// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAITaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAITaskStatusResponseBody
	GetAccountName() *string
	SetDBClusterId(v string) *DescribeAITaskStatusResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeAITaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAITaskStatusResponseBody
	GetStatus() *string
	SetStatusName(v string) *DescribeAITaskStatusResponseBody
	GetStatusName() *string
}

type DescribeAITaskStatusResponseBody struct {
	// The name of the database account that is used to connect to the AI nodes in the cluster.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the PolarDB for AI feature. Valid values:
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: disabled.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the status of the PolarDB for AI feature.
	//
	// example:
	//
	// Closed State
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
}

func (s DescribeAITaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAITaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAITaskStatusResponseBody) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAITaskStatusResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAITaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAITaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAITaskStatusResponseBody) GetStatusName() *string {
	return s.StatusName
}

func (s *DescribeAITaskStatusResponseBody) SetAccountName(v string) *DescribeAITaskStatusResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeAITaskStatusResponseBody) SetDBClusterId(v string) *DescribeAITaskStatusResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAITaskStatusResponseBody) SetRequestId(v string) *DescribeAITaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAITaskStatusResponseBody) SetStatus(v string) *DescribeAITaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAITaskStatusResponseBody) SetStatusName(v string) *DescribeAITaskStatusResponseBody {
	s.StatusName = &v
	return s
}

func (s *DescribeAITaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
