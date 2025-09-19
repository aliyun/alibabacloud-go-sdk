// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorTrialConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCloudVendorTrialConfigResponseBody
	GetRequestId() *string
}

type AddCloudVendorTrialConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCloudVendorTrialConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorTrialConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddCloudVendorTrialConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCloudVendorTrialConfigResponseBody) SetRequestId(v string) *AddCloudVendorTrialConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCloudVendorTrialConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
