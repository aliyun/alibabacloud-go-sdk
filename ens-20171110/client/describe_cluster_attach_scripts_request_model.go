// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttachScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterAttachScriptsRequest
	GetClusterId() *string
	SetNodepoolId(v string) *DescribeClusterAttachScriptsRequest
	GetNodepoolId() *string
	SetOptions(v string) *DescribeClusterAttachScriptsRequest
	GetOptions() *string
}

type DescribeClusterAttachScriptsRequest struct {
	// The ID of the ECK cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// eck-11111111
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the node pool to which the node will be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
	// Optional parameters for the script, provided as a JSON string.
	//
	// example:
	//
	// {\\"key1\\":\\"val1\\"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s DescribeClusterAttachScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttachScriptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterAttachScriptsRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterAttachScriptsRequest) GetOptions() *string {
	return s.Options
}

func (s *DescribeClusterAttachScriptsRequest) SetClusterId(v string) *DescribeClusterAttachScriptsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetNodepoolId(v string) *DescribeClusterAttachScriptsRequest {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetOptions(v string) *DescribeClusterAttachScriptsRequest {
	s.Options = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) Validate() error {
	return dara.Validate(s)
}
