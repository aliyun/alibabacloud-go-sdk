// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveDomainMultiStreamConfigResponseBody
	GetRequestId() *string
}

type SetLiveDomainMultiStreamConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveDomainMultiStreamConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveDomainMultiStreamConfigResponseBody) SetRequestId(v string) *SetLiveDomainMultiStreamConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveDomainMultiStreamConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
