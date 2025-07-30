// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReverseWriterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckPoint(v string) *StartReverseWriterRequest
	GetCheckPoint() *string
	SetDtsJobId(v string) *StartReverseWriterRequest
	GetDtsJobId() *string
	SetResourceGroupId(v string) *StartReverseWriterRequest
	GetResourceGroupId() *string
}

type StartReverseWriterRequest struct {
	// The offset of the Incremental Write module. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > The default value is the offset that is automatically saved by DTS when the task is paused.
	//
	// example:
	//
	// 1695613785
	CheckPoint *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	// The ID of the reverse task that was created by calling the CreateReverseDtsJob operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// n99m9jx822k****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StartReverseWriterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartReverseWriterRequest) GoString() string {
	return s.String()
}

func (s *StartReverseWriterRequest) GetCheckPoint() *string {
	return s.CheckPoint
}

func (s *StartReverseWriterRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *StartReverseWriterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartReverseWriterRequest) SetCheckPoint(v string) *StartReverseWriterRequest {
	s.CheckPoint = &v
	return s
}

func (s *StartReverseWriterRequest) SetDtsJobId(v string) *StartReverseWriterRequest {
	s.DtsJobId = &v
	return s
}

func (s *StartReverseWriterRequest) SetResourceGroupId(v string) *StartReverseWriterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartReverseWriterRequest) Validate() error {
	return dara.Validate(s)
}
