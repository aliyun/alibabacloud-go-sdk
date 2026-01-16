// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpgradeClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpgradeClusterAddonsResponseBody
	GetRequestId() *string
}

type UpgradeClusterAddonsResponseBody struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClusterAddonsResponseBody) SetClusterId(v string) *UpgradeClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterAddonsResponseBody) SetRequestId(v string) *UpgradeClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
