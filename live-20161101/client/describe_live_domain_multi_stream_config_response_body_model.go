// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMultiStreamConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveDomainMultiStreamConfigResponseBody
	GetRequestId() *string
	SetSwitch(v string) *DescribeLiveDomainMultiStreamConfigResponseBody
	GetSwitch() *string
}

type DescribeLiveDomainMultiStreamConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F486A44F-6B35-5A96-BF2C-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the dual-stream disaster recovery feature is enabled. Valid values:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// on
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
}

func (s DescribeLiveDomainMultiStreamConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMultiStreamConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMultiStreamConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainMultiStreamConfigResponseBody) GetSwitch() *string {
	return s.Switch
}

func (s *DescribeLiveDomainMultiStreamConfigResponseBody) SetRequestId(v string) *DescribeLiveDomainMultiStreamConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigResponseBody) SetSwitch(v string) *DescribeLiveDomainMultiStreamConfigResponseBody {
	s.Switch = &v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
