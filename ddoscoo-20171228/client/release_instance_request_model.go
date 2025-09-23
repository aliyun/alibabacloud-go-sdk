// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseInstanceRequest
	GetInstanceId() *string
	SetSourceIp(v string) *ReleaseInstanceRequest
	GetSourceIp() *string
}

type ReleaseInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseInstanceRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetSourceIp(v string) *ReleaseInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
