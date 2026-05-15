// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKbUuid(v string) *DescribeKnowledgeBaseStatsRequest
	GetKbUuid() *string
}

type DescribeKnowledgeBaseStatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DescribeKnowledgeBaseStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseStatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseStatsRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeKnowledgeBaseStatsRequest) SetKbUuid(v string) *DescribeKnowledgeBaseStatsRequest {
	s.KbUuid = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsRequest) Validate() error {
	return dara.Validate(s)
}
