// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *VerifyCheckResultRequest
	GetCheckIds() []*int64
	SetTaskSource(v string) *VerifyCheckResultRequest
	GetTaskSource() *string
}

type VerifyCheckResultRequest struct {
	// The IDs of the check items.
	CheckIds   []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	TaskSource *string  `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s VerifyCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckResultRequest) GoString() string {
	return s.String()
}

func (s *VerifyCheckResultRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *VerifyCheckResultRequest) GetTaskSource() *string {
	return s.TaskSource
}

func (s *VerifyCheckResultRequest) SetCheckIds(v []*int64) *VerifyCheckResultRequest {
	s.CheckIds = v
	return s
}

func (s *VerifyCheckResultRequest) SetTaskSource(v string) *VerifyCheckResultRequest {
	s.TaskSource = &v
	return s
}

func (s *VerifyCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
