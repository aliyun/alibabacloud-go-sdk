// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceWhiteListResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ConfigInstanceWhiteListResponseBody
	GetRequestId() *string
}

type ConfigInstanceWhiteListResponseBody struct {
	// The ID of the bastion host for which a whitelist of public IP addresses is configured.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 47820E32-5968-45CF-982F-09CB80DC180B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceWhiteListResponseBody) SetInstanceId(v string) *ConfigInstanceWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhiteListResponseBody) SetRequestId(v string) *ConfigInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
