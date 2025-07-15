// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamOptimalModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveDomainMultiStreamOptimalModeResponseBody
	GetRequestId() *string
}

type SetLiveDomainMultiStreamOptimalModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveDomainMultiStreamOptimalModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamOptimalModeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamOptimalModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveDomainMultiStreamOptimalModeResponseBody) SetRequestId(v string) *SetLiveDomainMultiStreamOptimalModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeResponseBody) Validate() error {
	return dara.Validate(s)
}
