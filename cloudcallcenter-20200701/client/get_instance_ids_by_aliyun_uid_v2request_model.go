// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdsByAliyunUidV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *GetInstanceIdsByAliyunUidV2Request
	GetAliyunUid() *string
}

type GetInstanceIdsByAliyunUidV2Request struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
}

func (s GetInstanceIdsByAliyunUidV2Request) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2Request) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2Request) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GetInstanceIdsByAliyunUidV2Request) SetAliyunUid(v string) *GetInstanceIdsByAliyunUidV2Request {
	s.AliyunUid = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Request) Validate() error {
	return dara.Validate(s)
}
