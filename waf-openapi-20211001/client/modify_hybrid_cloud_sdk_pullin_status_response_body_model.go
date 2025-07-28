// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudSdkPullinStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudSdkPullinStatusResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudSdkPullinStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B191**EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudSdkPullinStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudSdkPullinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudSdkPullinStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudSdkPullinStatusResponseBody) SetRequestId(v string) *ModifyHybridCloudSdkPullinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
