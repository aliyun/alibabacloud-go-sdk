// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSyncAlgorithmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *InvokeSyncAlgorithmShrinkRequest
	GetJobId() *string
	SetName(v string) *InvokeSyncAlgorithmShrinkRequest
	GetName() *string
	SetParamsShrink(v string) *InvokeSyncAlgorithmShrinkRequest
	GetParamsShrink() *string
}

type InvokeSyncAlgorithmShrinkRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s InvokeSyncAlgorithmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeSyncAlgorithmShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeSyncAlgorithmShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *InvokeSyncAlgorithmShrinkRequest) GetName() *string {
	return s.Name
}

func (s *InvokeSyncAlgorithmShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *InvokeSyncAlgorithmShrinkRequest) SetJobId(v string) *InvokeSyncAlgorithmShrinkRequest {
	s.JobId = &v
	return s
}

func (s *InvokeSyncAlgorithmShrinkRequest) SetName(v string) *InvokeSyncAlgorithmShrinkRequest {
	s.Name = &v
	return s
}

func (s *InvokeSyncAlgorithmShrinkRequest) SetParamsShrink(v string) *InvokeSyncAlgorithmShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *InvokeSyncAlgorithmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
