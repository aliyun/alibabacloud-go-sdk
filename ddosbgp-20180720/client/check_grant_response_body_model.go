// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckGrantResponseBody
	GetRequestId() *string
	SetStatus(v int32) *CheckGrantResponseBody
	GetStatus() *int32
}

type CheckGrantResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DB002CE5-5E6C-5F11-AE15-B525299A40F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account. Valid values:
	//
	// 	- **1**: Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
	//
	// 	- **0**: Anti-DDoS Origin is not authorized to obtain information about the assets within the current Alibaba Cloud account.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckGrantResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

func (s *CheckGrantResponseBody) Validate() error {
	return dara.Validate(s)
}
