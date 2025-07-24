// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPod interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *Pod
	GetMessage() *string
	SetPodName(v string) *Pod
	GetPodName() *string
	SetPodStatus(v string) *Pod
	GetPodStatus() *string
	SetReason(v string) *Pod
	GetReason() *string
}

type Pod struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PodName   *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	PodStatus *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s Pod) String() string {
	return dara.Prettify(s)
}

func (s Pod) GoString() string {
	return s.String()
}

func (s *Pod) GetMessage() *string {
	return s.Message
}

func (s *Pod) GetPodName() *string {
	return s.PodName
}

func (s *Pod) GetPodStatus() *string {
	return s.PodStatus
}

func (s *Pod) GetReason() *string {
	return s.Reason
}

func (s *Pod) SetMessage(v string) *Pod {
	s.Message = &v
	return s
}

func (s *Pod) SetPodName(v string) *Pod {
	s.PodName = &v
	return s
}

func (s *Pod) SetPodStatus(v string) *Pod {
	s.PodStatus = &v
	return s
}

func (s *Pod) SetReason(v string) *Pod {
	s.Reason = &v
	return s
}

func (s *Pod) Validate() error {
	return dara.Validate(s)
}
