// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdsByAliyunUidV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetHttpStatusCode() *int64
	SetInstanceIds(v []*string) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceIdsByAliyunUidV2ResponseBody
	GetSuccess() *bool
}

type GetInstanceIdsByAliyunUidV2ResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceIds    []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceIdsByAliyunUidV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetCode(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetHttpStatusCode(v int64) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetInstanceIds(v []*string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.InstanceIds = v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetMessage(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetRequestId(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetSuccess(v bool) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
