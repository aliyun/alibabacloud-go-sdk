// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamMasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveDomainMultiStreamMasterResponseBody
	GetRequestId() *string
}

type SetLiveDomainMultiStreamMasterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4E*****43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveDomainMultiStreamMasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamMasterResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamMasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveDomainMultiStreamMasterResponseBody) SetRequestId(v string) *SetLiveDomainMultiStreamMasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterResponseBody) Validate() error {
	return dara.Validate(s)
}
