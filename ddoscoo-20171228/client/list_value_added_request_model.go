// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValueAddedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ListValueAddedRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *ListValueAddedRequest
	GetSourceIp() *string
}

type ListValueAddedRequest struct {
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListValueAddedRequest) String() string {
	return dara.Prettify(s)
}

func (s ListValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ListValueAddedRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListValueAddedRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListValueAddedRequest) SetResourceGroupId(v string) *ListValueAddedRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListValueAddedRequest) SetSourceIp(v string) *ListValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ListValueAddedRequest) Validate() error {
	return dara.Validate(s)
}
