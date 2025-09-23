// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseValueAddedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseValueAddedRequest
	GetInstanceId() *string
	SetSourceIp(v string) *ReleaseValueAddedRequest
	GetSourceIp() *string
}

type ReleaseValueAddedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ddos_fl_pre-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ReleaseValueAddedRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseValueAddedRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ReleaseValueAddedRequest) SetInstanceId(v string) *ReleaseValueAddedRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseValueAddedRequest) SetSourceIp(v string) *ReleaseValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseValueAddedRequest) Validate() error {
	return dara.Validate(s)
}
