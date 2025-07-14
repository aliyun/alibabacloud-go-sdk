// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebInstanceLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebInstanceLogsRequest
	GetNamespaceId() *string
}

type DescribeWebInstanceLogsRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebInstanceLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceLogsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebInstanceLogsRequest) SetNamespaceId(v string) *DescribeWebInstanceLogsRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebInstanceLogsRequest) Validate() error {
	return dara.Validate(s)
}
