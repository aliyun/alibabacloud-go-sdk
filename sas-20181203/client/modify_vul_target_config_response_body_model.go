// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVulTargetConfigResponseBody
	GetRequestId() *string
}

type ModifyVulTargetConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1FF908BA-ADD8-5138-8595-614C6E3C6658
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVulTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVulTargetConfigResponseBody) SetRequestId(v string) *ModifyVulTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVulTargetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
